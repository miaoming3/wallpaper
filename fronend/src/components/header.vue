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
  <el-menu
      :default-active="props.activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      :ellipsis="false"
      router=true
  >
    <el-menu-item index="/admin">
      <h1>kmi壁纸</h1>
    </el-menu-item>
    <el-menu-item index="1">dd</el-menu-item>
    <el-sub-menu index="2">
      <template #title>admin</template>
      <el-menu-item index="2-1">我的资料</el-menu-item>
      <el-menu-item index="/admin/changePassword">修改密码</el-menu-item>
      <el-menu-item index="2-3">退出</el-menu-item>
    </el-sub-menu>
  </el-menu>
</template>

<style scoped lang="scss">


.el-menu--horizontal > .el-menu-item:nth-child(1) {
  margin-right: auto;
  &:hover,&:focus{
    background-color: white;
  }
  &.is-active{
    background:none;
    border-bottom: none;
  }
}
</style>