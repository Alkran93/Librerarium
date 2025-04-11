package src.gateway.util;

import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;

import java.io.IOException;
import java.io.OutputStream;

public class AuthFilter implements HttpHandler {

    private final HttpHandler next;

    public AuthFilter(HttpHandler next) {
        this.next = next;
    }

    @Override
    public void handle(HttpExchange exchange) throws IOException {
        String path = exchange.getRequestURI().getPath();
        LoggerUtil.log("Authenticating request to: " + path);

        String authHeader = exchange.getRequestHeaders().getFirst("Authorization");

        if (authHeader == null || !authHeader.startsWith("Bearer ")) {
            sendUnauthorized(exchange, "Missing or invalid Authorization header.");
            return;
        }

        String token = authHeader.substring(7);

        if (!JwtUtil.isValid(token)) {
            sendUnauthorized(exchange, "Invalid or expired token.");
            return;
        }

        // Token is valid, pass the request to the next handler
        next.handle(exchange);
    }

    private void sendUnauthorized(HttpExchange exchange, String message) throws IOException {
        String response = "{\"error\": \"" + message + "\"}";
        exchange.getResponseHeaders().add("Content-Type", "application/json");
        exchange.sendResponseHeaders(401, response.getBytes().length);

        try (OutputStream os = exchange.getResponseBody()) {
            os.write(response.getBytes());
        }

        LoggerUtil.log("Unauthorized access: " + message);
    }
}

