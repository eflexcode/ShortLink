package com.ifeanyi.url_server.service;

import com.ifeanyi.url_server.entity.UrlEntity;
import com.ifeanyi.url_server.exceptions.NotFoundException;
import com.ifeanyi.url_server.model.UrlPayload;
import com.ifeanyi.url_server.model.UrlResponse;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;

public interface UrlService {

    UrlResponse createShortLink(UrlPayload urlPayload) throws NotFoundException;
    Page<UrlEntity> getByOwnerId(String id, Pageable pageable);
    UrlEntity getByUrl(String url) throws NotFoundException;

}
