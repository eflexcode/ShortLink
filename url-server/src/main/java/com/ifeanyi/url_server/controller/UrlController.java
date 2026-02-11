package com.ifeanyi.url_server.controller;

import com.ifeanyi.url_server.entity.UrlEntity;
import com.ifeanyi.url_server.exceptions.NotFoundException;
import com.ifeanyi.url_server.model.UrlPayload;
import com.ifeanyi.url_server.model.UrlResponse;
import com.ifeanyi.url_server.service.UrlService;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.http.HttpStatus;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.io.IOException;
import java.util.List;

@RestController
@RequestMapping("/")
@RequiredArgsConstructor
public class UrlController {

    private final UrlService urlService;
    private String basePathToCutUrl = "http://localhost:8083/r/";

    @PostMapping("create-short-url")
    public UrlResponse createShortLink(@RequestBody UrlPayload urlPayload) throws NotFoundException {
        return urlService.createShortLink(urlPayload);
    }

    @GetMapping("r/**")
    public void getUrl(HttpServletRequest request, HttpServletResponse response) throws NotFoundException, IOException {

        String buffer = String.valueOf(request.getRequestURL());

        String path = buffer.replaceFirst(basePathToCutUrl, "");

        UrlEntity gottenUrl = urlService.getByUrl(path);
        response.sendRedirect(gottenUrl.getUrlOriginal());
    }

    @GetMapping("get/{id}")
    public List<UrlEntity> getByOwnerId(@PathVariable String id, Pageable pageable) {
        return urlService.getByOwnerId(id, pageable).stream().toList();
    }

}
