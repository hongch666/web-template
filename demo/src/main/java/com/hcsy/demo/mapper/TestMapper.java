package com.hcsy.demo.mapper;

import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;

@Mapper
public interface TestMapper {
    @Select("SELECT 'Hello, Mybatis-Plus!' ")
    String test();
}
