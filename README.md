# tdd

[![Build Status](https://www.travis-ci.org/63isOK/tdd.svg?branch=master)](https://www.travis-ci.org/63isOK/tdd)

练习tdd流程的仓库,附加ci.
因为是一个练习项目,所以不走pr

## 依赖注入

di,ioc的一种,用途是解耦,通常是用接口类型来代替具体类型

下面几种场景才会体会到di的威力:
- 测试代码
- 关注分离点 某某在何时做了何事 完全可以分离,测试方便,复用性也高
- 代码复用

## 迭代

什么叫迭代:
- 用最小的步骤,让软件可用
- 一个迭代,正好对应一个tdd循环
- tdd是xp(极限编程)的一种,xp属于敏捷的一种,迭代又术语敏捷的核心词汇
- 所以一次迭代用tdd的 红灯-绿灯-重构 来解释,刚刚好,都强调了最小步骤达成
- 写的程序在硬盘挂掉之后还能跑,这种做法和敏捷一点关系都没有

## kent beck

- 是tdd和极限编程的开创者
- 所有编码的策略都应该是:
    - make it work 对应tdd的绿灯,最小满足需求的实现 
    - make it right 对应tdd的重构
    - make it fast
- 过早的优化是万恶之源,优化不要在绿灯/红灯阶段做,要放在重构中

## road map

- [ ] 基础tdd练习
- [ ] ci
