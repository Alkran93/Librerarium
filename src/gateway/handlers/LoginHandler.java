package src.gateway.handlers;

import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;
import src.gateway.util.Config;
import src.gateway.util.LoggerUtil;

import java.io.*;
import java.net.HttpURLConnection;
import java.net.URL;
import java.util.stream.Collectors;

public class LoginHandler implements HttpHandler {
    @Override
    public void handle(HttpExchange exchange) {
        try {
            LoggerUtil.log("Received request to /login: " + exchange.getRequestMethod());

            URL url = new URL(Config.USER_SERVICE_URL + "/login");
            HttpURLConnection conn = (HttpURLConnection) url.openConnection();
            conn.setRequestMethod("POST");
            conn.setDoOutput(true);
            conn.setRequestProperty("Content-Type", "application/json");

            // Forward client body to auth service
            String requestBody = new BufferedReader(new InputStreamReader(exchange.getRequestBody()))
                    .lines().collect(Collectors.joining("\n"));

            try (OutputStream os = conn.getOutputStream()) {
                os.write(requestBody.getBytes());
            }

            int responseCode = conn.getResponseCode();
            LoggerUtil.log("Forwarded to user service. Response code: " + responseCode);

            InputStream responseStream = responseCode >= 200 && responseCode < 300
                    ? conn.getInputStream()
                    : conn.getErrorStream();

            String responseBody = new BufferedReader(new InputStreamReader(responseStream))
                    .lines().collect(Collectors.joining("\n"));

            exchange.getResponseHeaders().set("Content-Type", "application/json");
            exchange.sendResponseHeaders(responseCode, responseBody.getBytes().length);

            try (OutputStream os = exchange.getResponseBody()) {
                os.write(responseBody.getBytes());
            }

        } catch (Exception e) {
            LoggerUtil.log("Error handling /login: " + e.getMessage());
            sendJsonError(exchange, 503, "User authentication service is unavailable.");
        }
    }

    private void sendJsonError(HttpExchange exchange, int statusCode, String message) {
        try {
            String json = "{ \"error\": \"" + message + "\" }";
            exchange.getResponseHeaders().add("Content-Type", "application/json");
            exchange.sendResponseHeaders(statusCode, json.getBytes().length);
            try (OutputStream os = exchange.getResponseBody()) {
                os.write(json.getBytes());
            }
        } catch (Exception ex) {
            LoggerUtil.log("Failed to send error response: " + ex.getMessage());
        }
    }
}
