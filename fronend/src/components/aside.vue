<script setup lang="ts">
import { PropType } from 'vue';

const props = defineProps({
  activeIndex: {
    type: String,
    required: true
  }
});
</script>

<template>
  <el-aside width="200px">
    <el-menu
        :default-active="props.activeIndex"
        class="el-menu-vertical-demo"
        router=true
    >
      <el-menu-item index="/admin">
        <el-icon><DataBoard /></el-icon>
        <span>数据面板</span>
      </el-menu-item>
      <el-sub-menu index="/menu">
        <template #title>
          <el-icon><Menu /></el-icon>
          <span>菜单管理</span>
        </template>
        <el-menu-item index="/admin/menu/index">菜单列表</el-menu-item>
        <el-menu-item index="/admin/menu/rule">权限列表</el-menu-item>
        <el-menu-item index="/admin/menu/role">角色列表</el-menu-item>
      </el-sub-menu>
      <el-sub-menu index="/user">
        <template #title>
          <el-icon><User /></el-icon>
          <span>用户管理</span>
        </template>
        <el-menu-item index="/admin/user/index">用户列表</el-menu-item>
        <el-menu-item index="/admin/user/grade">用户等级</el-menu-item>
      </el-sub-menu>
      <el-sub-menu index="/administrators">
        <template #title>
          <el-icon><UserFilled /></el-icon>
          <span>管理员管理</span>
        </template>
        <el-menu-item index="/admin/index">管理员列表</el-menu-item>
      </el-sub-menu>
      <el-sub-menu index="/category">
        <template #title>
          <el-icon><Document /></el-icon>
          <span>分类管理</span>
        </template>
        <el-menu-item index="/admin/category/index">分类列表</el-menu-item>
      </el-sub-menu>
      <el-sub-menu index="/Picture">
        <template #title>
          <el-icon><Picture /></el-icon>
          <span>图片管理</span>
        </template>
        <el-menu-item index="/admin/img/index">图片列表</el-menu-item>
        <el-menu-item index="/admin/img/tags">标签列表</el-menu-item>
      </el-sub-menu>
      <el-sub-menu index="/setting">
        <template #title>
          <el-icon><Setting /></el-icon>
          <span>系统设置</span>
        </template>
        <el-menu-item index="/admin/setting">基本设置</el-menu-item>
      </el-sub-menu>
    </el-menu>

  </el-aside>
</template>

<style scoped lang="scss">
.el-menu{
  height: 100%;
}
</style>