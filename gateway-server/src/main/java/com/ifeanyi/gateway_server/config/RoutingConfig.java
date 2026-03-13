package com.ifeanyi.gateway_server.config;

import org.springframework.cloud.gateway.route.RouteLocator;
import org.springframework.cloud.gateway.route.builder.RouteLocatorBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class RoutingConfig {

    @Bean
    public RouteLocator routeLocator(RouteLocatorBuilder builder){

        return builder.routes()
                .route("url", r -> r
                        .path("/url/**")
                        .filters(f -> f.stripPrefix(1))
                        .uri("http://localhost:8083"))
                .route("auth", r -> r
                        .path("/auth/**")
                        .filters(f -> f.stripPrefix(1))
                        .uri("http://localhost:8084"))
                .route("user", r -> r
                        .path("/user/**")
                        .filters(f -> f.stripPrefix(1))
                        .uri("http://localhost:8082"))
                .build();

    }

}
