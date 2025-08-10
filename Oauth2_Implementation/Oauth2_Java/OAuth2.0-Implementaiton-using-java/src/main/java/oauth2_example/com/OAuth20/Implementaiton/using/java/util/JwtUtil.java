package oauth2_example.com.OAuth20.Implementaiton.using.java.util;

import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import io.jsonwebtoken.security.Keys;

import java.security.Key;
import java.util.Date;
import java.util.Map;

public class JwtUtil {
    private final Key key;
    private final long expirationSeconds;

    public JwtUtil(String baseSecret, long expirationSeconds){
        this.key= Keys.hmacShaKeyFor(baseSecret.getBytes());
        this.expirationSeconds=expirationSeconds;
    }
    public String createToken(Map<String, Object> claims){
        Date now=new Date();
        return Jwts.builder().setClaims(claims).setIssuedAt(now).setExpiration(new Date(now.getTime()+expirationSeconds*1000))
                .signWith(key, SignatureAlgorithm.ES256).compact();
    }
}
