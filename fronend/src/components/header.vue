<script setup lang="ts">
import {Logout} from "@/assets/api/admin";
import {ElMessage} from "element-plus";
import router from "@/router";
const props = defineProps({
  activeIndex: {
    type: String,
    required: true
  },
  adminInfo: {
    type: Object,
    required: true,
  },
});

const onLogout=async () =>{
  try {
    const res= await  Logout();
    if (res.code===0) {
      localStorage.clear();
      ElMessage.success({message: '退出成功,即将重定向', duration: 1000});
      setTimeout(() => {
        router.replace({path: '/'});
      }, 2000);
    }
  }catch(e){
  }

}


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
    <el-sub-menu index="adminInfo">
      <template #title>{{ props.adminInfo.username }}</template>
      <el-menu-item index="/admin/info">我的资料</el-menu-item>
      <el-menu-item index="/admin/changePassword">修改密码</el-menu-item>
      <el-menu-item  @click="onLogout">退出</el-menu-item>
    </el-sub-menu>
  </el-menu>

</template>

<style scoped lang="scss">


.el-menu--horizontal > .el-menu-item:nth-child(1) {
  margin-right: auto;

  &:hover, &:focus {
    background-color: white;
  }

  &.is-active {
    background: none;
    border-bottom: none;
  }
}
</style>
