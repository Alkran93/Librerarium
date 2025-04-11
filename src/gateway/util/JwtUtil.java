package src.gateway.util;

import java.util.Base64;

public class JwtUtil {
    private static final String SECRET = "super-secret-key";

    public static String createToken(String username) {
        String header = "{\"alg\":\"HS256\",\"typ\":\"JWT\"}";
        String payload = "{\"sub\":\"" + username + "\",\"role\":\"user\"}";

        String encodedHeader = base64UrlEncode(header.getBytes());
        String encodedPayload = base64UrlEncode(payload.getBytes());

        String data = encodedHeader + "." + encodedPayload;
        String signature = base64UrlEncode(hmacSha256(data, SECRET));

        return data + "." + signature;
    }

    public static boolean isValid(String token) {
        try {
            String[] parts = token.split("\\.");
            if (parts.length != 3) {
                System.out.println("Token format invalid");
                return false;
            }

            String header = parts[0];
            String payload = parts[1];
            String signature = parts[2];

            String data = header + "." + payload;
            String expectedSignature = base64UrlEncode(hmacSha256(data, SECRET));

            return expectedSignature.equals(signature);
        } catch (Exception e) {
            System.out.println("Token validation error: " + e.getMessage());
            return false;
        }
    }

    private static byte[] hmacSha256(String data, String key) {
        try {
            javax.crypto.Mac mac = javax.crypto.Mac.getInstance("HmacSHA256");
            javax.crypto.spec.SecretKeySpec secretKey = new javax.crypto.spec.SecretKeySpec(key.getBytes(), "HmacSHA256");
            mac.init(secretKey);
            return mac.doFinal(data.getBytes());
        } catch (Exception e) {
            throw new RuntimeException("HMAC SHA256 Error", e);
        }
    }

    private static String base64UrlEncode(byte[] input) {
        return Base64.getUrlEncoder().withoutPadding().encodeToString(input);
    }
}
