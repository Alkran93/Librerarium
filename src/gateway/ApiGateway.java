package src.gateway;


import com.sun.net.httpserver.HttpServer;

import src.gateway.handlers.*;
import src.gateway.util.AuthFilter;

import java.io.IOException;
import java.net.InetSocketAddress;

public class ApiGateway {
    public static void main(String[] args) throws IOException {
        int port = 8080;
        HttpServer server = HttpServer.create(new InetSocketAddress(port), 0);

        server.createContext("/", new RootHandler());

        // Public endpoint (no token required)
        server.createContext("/login", new LoginHandler());

        // Secured endpoints (require valid JWT)
        server.createContext("/products", new AuthFilter(new ProductsHandler()));
        server.createContext("/cart", new AuthFilter(new CartHandler()));

        System.out.println("API Gateway running on port " + port);
        server.start();
    }
}
