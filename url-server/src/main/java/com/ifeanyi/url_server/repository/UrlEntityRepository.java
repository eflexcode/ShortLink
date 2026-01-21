package com.ifeanyi.url_server.repository;

import com.ifeanyi.url_server.entity.UrlEntity;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface UrlEntityRepository extends MongoRepository<UrlEntity,String> {
    Page<UrlEntity> findByOwnerId(String ownerId, Pageable pageable);
    Optional<UrlEntity> findByUrlShort(String urlShort);
}
