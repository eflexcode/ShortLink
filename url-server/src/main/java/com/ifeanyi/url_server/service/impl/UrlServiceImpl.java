package com.ifeanyi.url_server.service.impl;

import com.ifeanyi.url_server.entity.UrlEntity;
import com.ifeanyi.url_server.exceptions.NotFoundException;
import com.ifeanyi.url_server.model.UrlPayload;
import com.ifeanyi.url_server.model.UrlResponse;
import com.ifeanyi.url_server.repository.UrlEntityRepository;
import com.ifeanyi.url_server.service.UrlService;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.BeanUtils;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;

import java.math.BigInteger;
import java.util.Date;
import java.util.Random;

@Service
@RequiredArgsConstructor
public class UrlServiceImpl implements UrlService {

    private UrlEntityRepository urlEntityRepository;

    @Override
    public UrlResponse createShortLink(UrlPayload urlPayload) {

        String shortUrl = shortUrl(urlPayload.getUrlOriginal());

        UrlEntity urlEntity = new UrlEntity();
        BeanUtils.copyProperties(urlPayload, urlEntity);
        urlEntity.setHits(BigInteger.ZERO);
        urlEntity.setUrlShort(shortUrl);
        urlEntity.setCreatedAt(new Date());

        UrlEntity saveUrlEntity = urlEntityRepository.save(urlEntity);

        shortUrl = saveUrlEntity.getUrlShort();

        return new UrlResponse(shortUrl);
    }

    @Override
    public Page<UrlEntity> getByOwnerId(String id, Pageable pageable) {
        return urlEntityRepository.findByOwnerId(id, pageable);
    }

    @Override
    public UrlEntity getByUrl(String url) throws NotFoundException {

        UrlEntity urlEntity = urlEntityRepository.findByUrl(url).orElseThrow(() -> new NotFoundException("no url found"));

        BigInteger newHitCount = urlEntity.getHits().add(BigInteger.ONE);
        urlEntity.setHits(newHitCount);
        urlEntityRepository.save(urlEntity);

        return urlEntity;
    }

    private String shortUrl(String url) {
// get first cha then last then any 2 random in length pus a dot
        return  url.substring(0, 0) + url.substring(url.length() - 1, url.length() - 1) + url.substring(new Random().nextInt(0, url.length() - 1))+ url.substring(new Random().nextInt(0, url.length() - 1))+".";

    }
}
