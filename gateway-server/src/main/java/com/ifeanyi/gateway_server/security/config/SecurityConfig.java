package com.ifeanyi.gateway_server.security.config;

import com.ifeanyi.gateway_server.security.filter.SecurityFilter;
import jakarta.servlet.Filter;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Lazy;
import org.springframework.security.authentication.AuthenticationProvider;
import org.springframework.security.authentication.dao.DaoAuthenticationProvider;
import org.springframework.security.config.annotation.method.configuration.EnableMethodSecurity;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configurers.AbstractHttpConfigurer;
import org.springframework.security.config.http.SessionCreationPolicy;
import org.springframework.security.web.SecurityFilterChain;
import org.springframework.security.web.authentication.AuthenticationFilter;
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter;

import java.util.AbstractCollection;

@Configuration
@EnableWebSecurity
@RequiredArgsConstructor
@EnableMethodSecurity
public class SecurityConfig {

    @Autowired
    @Lazy
    SecurityFilter filter;

    @Bean
    public SecurityFilterChain filterChain(HttpSecurity httpSecurity){

        httpSecurity.csrf(AbstractHttpConfigurer::disable)
                .sessionManagement(m ->   m.sessionCreationPolicy(SessionCreationPolicy.STATELESS))
                .authorizeHttpRequests(r  -> r.requestMatchers("/auth/**")
                        .permitAll()
                        .anyRequest()
                        .authenticated())
                .addFilterBefore(filter, UsernamePasswordAuthenticationFilter.class);

        return httpSecurity.build();
    }

}
