package com.ifeanyi.gateway_server.security.filter;

import com.ifeanyi.gateway_server.util.Util;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import jakarta.servlet.FilterChain;
import jakarta.servlet.ServletException;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationProvider;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;
import org.springframework.web.server.ResponseStatusException;

import java.io.IOException;
import java.util.Date;
import java.util.function.Function;

@Component
@EnableWebSecurity
//@RequiredArgsConstructor
public class SecurityFilter extends OncePerRequestFilter {

    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain) throws ServletException, IOException {

        String header = request.getHeader("Authorization");

        if (header != null && header.startsWith("Bearer ")) {

            String token = header.substring(7);

            if(!verifyToken(token)){
                throw new ResponseStatusException(HttpStatus.UNAUTHORIZED,"Auth failed");
            }

        }

        filterChain.doFilter(request, response);

    }

    private boolean verifyToken(String token) {

        Date date = extractClaims(token,Claims::getExpiration);

        return new Date().after(date);
    }

    public  <T>  T extractClaims(String token,Function<Claims,T> tFunction){

        Claims claims = Jwts.parserBuilder()
                .setSigningKey(Util.TOKEN_KEY)
                .build().parseClaimsJws(token)
                .getBody();

        return tFunction.apply(claims);

    }

}

