package src.gateway.handlers;

import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;

import src.gateway.util.Config;
import src.gateway.util.LoggerUtil;

import java.io.InputStream;
import java.io.OutputStream;
import java.net.HttpURLConnection;
import java.net.URL;

public class ProductsHandler implements HttpHandler {
    @Override
    public void handle(HttpExchange exchange) {
        try {
            LoggerUtil.log("Received request to /products: " + exchange.getRequestMethod());

            URL url = new URL(Config.PRODUCT_SERVICE_URL + "/products");
            HttpURLConnection conn = (HttpURLConnection) url.openConnection();
            conn.setRequestMethod(exchange.getRequestMethod());
            conn.setDoOutput(true);
            conn.setRequestProperty("Content-Type", "application/json");

            if (exchange.getRequestMethod().equals("POST") || exchange.getRequestMethod().equals("PUT")) {
                try (InputStream is = exchange.getRequestBody();
                     OutputStream os = conn.getOutputStream()) {
                    is.transferTo(os);
                }
            }

            int responseCode = conn.getResponseCode();
            LoggerUtil.log("Forwarded to product service. Response code: " + responseCode);
            exchange.sendResponseHeaders(responseCode, conn.getInputStream().available());

            try (InputStream is = conn.getInputStream();
                 OutputStream os = exchange.getResponseBody()) {
                is.transferTo(os);
            }

        } catch (Exception e) {
            LoggerUtil.log("Error handling /products: " + e.getMessage());
            sendJsonError(exchange, 503, "Product service is unavailable.");
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
