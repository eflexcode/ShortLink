package com.ifeanyi.gateway_server.security.config;

import com.ifeanyi.gateway_server.security.filter.SecurityFilter;
import com.ifeanyi.gateway_server.user.repository.UserRepository;
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
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.web.SecurityFilterChain;
import org.springframework.security.web.authentication.AuthenticationFilter;
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter;

import java.util.AbstractCollection;

@Configuration
@EnableMethodSecurity
@EnableWebSecurity
@RequiredArgsConstructor
public class SecurityConfig {

    @Autowired
    @Lazy
    SecurityFilter filter;

    private final UserRepository repository;

    @Bean
    public UserDetailsService userDetailsService() {
        return username -> repository.findByUsername(username).orElseThrow(() -> new UsernameNotFoundException("No user found"));
    }

    @Bean
    public AuthenticationProvider authenticationProvider(){

        DaoAuthenticationProvider provider = new DaoAuthenticationProvider(userDetailsService());
        provider.setPasswordEncoder(new BCryptPasswordEncoder());

        return provider;
    }

    @Bean
    public SecurityFilterChain filterChain(HttpSecurity httpSecurity) throws Exception{

        httpSecurity
                .csrf(AbstractHttpConfigurer::disable)
                .sessionManagement(m -> m.sessionCreationPolicy(SessionCreationPolicy.STATELESS))
                .authorizeHttpRequests(r -> r.requestMatchers("/auth/**")
                        .permitAll()
                        .anyRequest()
                        .authenticated())
                .addFilterBefore(filter, UsernamePasswordAuthenticationFilter.class)
                .authenticationProvider(authenticationProvider());

        return httpSecurity.build();
    }

}

