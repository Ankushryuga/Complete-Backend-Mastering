package oauth2_example.com.OAuth20.Implementaiton.using.java.controller;

import io.netty.util.Timeout;
import jakarta.annotation.PostConstruct;
import oauth2_example.com.OAuth20.Implementaiton.using.java.util.JwtUtil;
import oauth2_example.com.OAuth20.Implementaiton.using.java.util.PKCEUtil;
import org.apache.commons.lang3.RandomStringUtils;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.util.LinkedMultiValueMap;
import org.springframework.util.MultiValueMap;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.reactive.function.client.WebClient;

import java.time.Duration;
import java.util.Map;
import java.util.Random;
import java.util.concurrent.TimeUnit;

@RestController
public class OAuthController {
    private final RedisTemplate<String, String> redisTemplate;
    private final WebClient webClient=WebClient.create();

    @Value("${app.oauth.client-id}")
    private String clientId;

    @Value("${app.oauth.client-secret}")
    private String clientSecret;

    @Value("${app.oauth.redirect-uri}")
    private String redirectUri;

    @Value("${app.oauth.scopes}")
    private String scopes;

    @Value("${app.jwt.secret}")
    private String jwtSecret;

    @Value("${app.jwt.expiration-seconds}")
    private long jwtExp;

    private JwtUtil jwtUtil;

    @PostConstruct
    public void init(){
        jwtUtil=new JwtUtil(jwtSecret, jwtExp);
    }

    @GetMapping("/login")
    public ResponseEntity<Void> login(){
        //create state and PKCE Verifier:
        String state= RandomStringUtils.randomAlphanumeric(32);
        String codeVerifier= PKCEUtil.generateCodeVerifier(64);
        String codeChallenge=PKCEUtil.generateCodeChallenge(codeVerifier);

        //Store verifier in redis keyed by state.
        redisTemplate.opsForValue().set(state, codeVerifier, 5, TimeUnit.MILLISECONDS);

        String authUrl = "https://accounts.google.com/o/oauth2/v2/auth"+
                "?response_type=code"+
                "&client_id="+clientId+
                "&redirect_uri="+redirectUri+
                "&scope="+scopes+
                "&state="+state+
                "&code_challenge="+codeChallenge+
                "&code_challenge_method=S256";
        HttpHeaders headers=new HttpHeaders();
        headers.setLocation(java.net.URI.create(authUrl));
        return ResponseEntity.status(302).headers(headers).build();
    }

    @GetMapping("/callback")
    public ResponseEntity<Map<String, Object>> callback(@RequestParam("code") String code, @RequestParam("state") String state){
        //retrieve verifier.
        String codeVerifier=redisTemplate.opsForValue().get(state);
        if(codeVerifier==null){
            return ResponseEntity.status(401).body(Map.of("error", "Invalid or expired state"));
        }

        //exchange code for token.
        //MultiValue or MultiValued is special type of map that allow multiple value associate to single key.
        MultiValueMap<String, String> form=new LinkedMultiValueMap<>();
        form.add("grant-type", "authorization_code");
        form.add("code", code);
        form.add("client_id", clientId);
        form.add("client_secret", clientSecret);
        form.add("redirect_uri", redirectUri);
        form.add("code_verifier", codeVerifier);

        Map tokenResp=webClient.post()
                .uri("https://oauth2.googleapis.com/token")
                .contentType(MediaType.APPLICATION_FORM_URLENCODED)
                .bodyValue(form)
                .retrieve()
                .bodyToMono(Map.class)
                .block(Duration.ofSeconds(10));

        if (tokenResp==null || tokenResp.get("access_token")==null){
            return ResponseEntity.status(500).body(Map.of("error", "Token exchange failed"));
        }

        String accessToken=tokenResp.get("access_token").toString();
        
    }

}
