package com.ifeanyi.url_server.service.impl;

import com.ifeanyi.url_server.Util;
import com.ifeanyi.url_server.entity.UrlEntity;
import com.ifeanyi.url_server.exceptions.NotFoundException;
import com.ifeanyi.url_server.model.UrlPayload;
import com.ifeanyi.url_server.model.UrlResponse;
import com.ifeanyi.url_server.model.User;
import com.ifeanyi.url_server.repository.UrlEntityRepository;
import com.ifeanyi.url_server.service.UrlService;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.BeanUtils;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.http.*;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestClientException;
import org.springframework.web.client.RestTemplate;

import java.util.Date;
import java.util.Random;

@Service
@RequiredArgsConstructor
public class UrlServiceImpl implements UrlService {

    private final UrlEntityRepository urlEntityRepository;
    private final RestTemplate restTemplate;

    @Override
    public UrlResponse createShortLink(UrlPayload urlPayload) throws NotFoundException {
        String shortUrl = shortUrl(urlPayload.getUrlOriginal());

        UrlEntity urlEntity = new UrlEntity();

        if (urlPayload.getOwnerId() != null) {

            try {
                ResponseEntity<User> userResponseEntity = restTemplate.exchange(Util.USER_SERVICE_HOST + "/v1/get-with-id/" + urlPayload.getOwnerId(), HttpMethod.GET, HttpEntity.EMPTY, User.class);
                if (userResponseEntity.getStatusCode().value() != HttpStatus.OK.value()) {
                    throw new NotFoundException("No user found with id: "+urlPayload.getOwnerId());
                }
            }catch (RestClientException e){
                throw new NotFoundException("No user found with id: "+urlPayload.getOwnerId());
            }


        }

        BeanUtils.copyProperties(urlPayload, urlEntity);
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

        UrlEntity urlEntity = urlEntityRepository.findByUrlShort(url).orElseThrow(() -> new NotFoundException("no url found"));

        long newHitCount = urlEntity.getHits() + 1L;
        urlEntity.setHits(newHitCount);
        urlEntityRepository.save(urlEntity);

        return urlEntity;
    }

    private String shortUrl(String url) {
        // get random cha then last then any 2 random in length pus a dot
        return getChar(url) + getChar(url) + getChar(url) + getChar(url) + ".";
    }

    private String getChar(String s) {
        String randChar = String.valueOf(s.charAt(new Random().nextInt(0, s.length() - 1)));
        if (randChar.equals("/")) {
            randChar = getChar(s);
        }
        return randChar;
    }

}
