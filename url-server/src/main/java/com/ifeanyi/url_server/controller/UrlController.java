package com.ifeanyi.url_server.controller;

import com.ifeanyi.url_server.entity.UrlEntity;
import com.ifeanyi.url_server.exceptions.NotFoundException;
import com.ifeanyi.url_server.model.UrlPayload;
import com.ifeanyi.url_server.model.UrlResponse;
import com.ifeanyi.url_server.service.UrlService;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/")
@RequiredArgsConstructor
public class UrlController {

    private final UrlService urlService;

    @PostMapping("create")
    public UrlResponse createShortLink(@RequestBody UrlPayload urlPayload) {
        return urlService.createShortLink(urlPayload);
    }

    @GetMapping("{url}")
    public String getUrl(@PathVariable String url) throws NotFoundException {
        UrlEntity gottenUrl = urlService.getByUrl(url);
        return "redirect:" + gottenUrl.getUrlOriginal();
    }

    @GetMapping("{id}")
    public Page<UrlEntity> getByOwnerId(@PathVariable String id, Pageable pageable) {
        return urlService.getByOwnerId(id, pageable);
    }


}
