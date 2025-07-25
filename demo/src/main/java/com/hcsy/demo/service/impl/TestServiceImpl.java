package com.hcsy.demo.service.impl;

import com.hcsy.demo.mapper.TestMapper;
import com.hcsy.demo.service.TestService;
import org.springframework.stereotype.Service;
import javax.annotation.Resource;

@Service
public class TestServiceImpl implements TestService {
    @Resource
    private TestMapper testMapper;

    @Override
    public String test() {
        return testMapper.test();
    }
}
