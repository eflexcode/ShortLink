package com.ifeanyi.url_server;

import com.ifeanyi.url_server.entity.UrlEntity;
import com.ifeanyi.url_server.model.UrlPayload;
import com.ifeanyi.url_server.model.UrlResponse;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.*;
import org.springframework.web.client.RestTemplate;

import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;

@SpringBootTest
class UrlServerApplicationTests {


    static RestTemplate restTemplate;
    String baseUrl = "http://localhost:8083";


    @BeforeAll
    static void init() {
        restTemplate = new RestTemplate();
    }


    @Test
    void contextLoads() {

    }

    @Test
    void testCreateEndpoint() {

        UrlPayload urlPayload = new UrlPayload();
        urlPayload.setOwnerId("ooo");
        urlPayload.setUrlOriginal("https://medium.com/@vivekrajyaguru1993/how-to-build-an-api-gateway-with-spring-cloud-gateway-and-eureka-the-beginners-guide-0985f0c42527");

        HttpEntity<UrlPayload> httpEntity = new HttpEntity<>(urlPayload);

        ResponseEntity<UrlResponse> exchange = restTemplate.exchange("http://localhost:8083/create-short-url", HttpMethod.POST, httpEntity, UrlResponse.class);
        assertEquals(exchange.getStatusCode().value(), HttpStatus.OK.value());

    }

    @Test
    void  testGetUrl(){

        HttpHeaders httpHeaders = new HttpHeaders();
        HttpEntity<HttpHeaders> httpEntity = new HttpEntity<>(httpHeaders);

        ResponseEntity<String> exchange = restTemplate.exchange("http://localhost:8083/r/vytc.", HttpMethod.GET, httpEntity, String.class);
        assertEquals(exchange.getStatusCode().value(), 302);
    }

    @Test
    void testBetById(){
        HttpHeaders httpHeaders = new HttpHeaders();
        HttpEntity<HttpHeaders> httpEntity = new HttpEntity<>(httpHeaders);

        ResponseEntity<List> exchange = restTemplate.exchange("http://localhost:8083/get/50392ebb-8ac1-4cce-9ce2-e4f9d84f9f2c", HttpMethod.GET, httpEntity,List.class);
        assertEquals(exchange.getStatusCode().value(), 200);
    }

}
