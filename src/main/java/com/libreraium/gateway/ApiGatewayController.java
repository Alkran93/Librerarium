package com.libreraium.gateway;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class ApiGatewayController {

    // Ruta de ejemplo para el microservicio de usuarios
    @GetMapping("/usuarios")
    public String getUsuarios() {
        return "Interacción con el microservicio de Usuarios";
    }

    // Ruta de ejemplo para el microservicio de productos
    @GetMapping("/productos")
    public String getProductos() {
        return "Interacción con el microservicio de Productos";
    }

    // Ruta de ejemplo para el microservicio de carrito
    @GetMapping("/carrito")
    public String getCarrito() {
        return "Interacción con el microservicio de Carrito";
    }

    // Ruta de ejemplo para el microservicio de pedidos
    @GetMapping("/pedidos")
    public String getPedidos() {
        return "Interacción con el microservicio de Pedidos";
    }
}
