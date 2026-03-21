package com.ifeanyi.url_server;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cache.annotation.EnableCaching;

@SpringBootApplication
@EnableCaching
public class UrlServerApplication {

	public static void main(String[] args) {
		SpringApplication.run(UrlServerApplication.class, args);
	}

}
