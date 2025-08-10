package oauth2_example.com.OAuth20.Implementaiton.using.java.util;

import java.security.MessageDigest;
import java.security.SecureRandom;
import java.util.Base64;

public class PKCEUtil {
    private static final SecureRandom secureRandom=new SecureRandom();
    public static String generateCodeVerifier(int length){
        byte[] bytes=new byte[length];
        secureRandom.nextBytes(bytes);
        return Base64.getUrlEncoder().withoutPadding().encodeToString(bytes);
    }

    public static String generateCodeChallenge(String verifier){
        try{
            MessageDigest md=MessageDigest.getInstance("SHA-256");
            byte[] digest=md.digest(verifier.getBytes("US-ASCII"));
            return Base64.getUrlEncoder().withoutPadding().encodeToString(digest);
        }catch (Exception e){
            throw new RuntimeException(e);
        }
    }
}
