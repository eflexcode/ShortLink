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
//import org.springframework.security.authentication.AuthenticationProvider;
//import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
//import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
//import org.springframework.security.core.context.SecurityContextHolder;
//import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;
import org.springframework.web.server.ResponseStatusException;
import org.springframework.web.server.ServerWebExchange;
import org.springframework.web.server.WebFilter;
import org.springframework.web.server.WebFilterChain;
import reactor.core.publisher.Mono;

import java.io.IOException;
import java.util.Date;
import java.util.function.Function;

@Component
public class SecurityFilter implements WebFilter {

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

    @Override
    public Mono<Void> filter(ServerWebExchange exchange, WebFilterChain chain) {

        String header = exchange.getRequest().getHeaders().getFirst("Authorization");

        if (header != null && header.startsWith("Bearer ")) {

            String token = header.substring(7);

            if(!verifyToken(token)){
                throw new ResponseStatusException(HttpStatus.UNAUTHORIZED,"Auth failed");
            }

        }

        return chain.filter(exchange);
    }

}

