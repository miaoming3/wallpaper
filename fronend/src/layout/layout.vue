<script setup lang="ts">
import {onMounted, provide, reactive, ref, watch} from 'vue'
import { useRouter } from 'vue-router';
import Aside from "@/components/aside.vue";
import Header from "@/components/header.vue";
import Footer from "@/components/footer.vue";
import {getAdminInfo} from "@/assets/api/admin";
const chooseHeader = window.location.pathname;
// 创建一个响应式引用并初始化为当前页面的路径
const activeIndex = ref<string>(chooseHeader);
const router = useRouter();
const adminInfo =reactive({
  username:"",
  email:"",
  phone:"",
  status:"",
  avatar:"",
  created_at:"",
  img_url:""
})
provide("adminInfo",adminInfo)

// 监听路由变化并更新 activeIndex
onMounted(async () => {
  watch(() => router.currentRoute.value.path,
      (newPath) => {activeIndex.value = newPath;});
  try {
    const data = await getAdminInfo();
    if (data && data.data) {
      Object.assign(adminInfo,data.data)
    }
  } catch (error) {
    console.error("Failed to fetch admin info:", error);
  }
});
</script>
<template>
  <div class="common-layout">
    <el-container>
      <el-header>
      <Header :admin-info="adminInfo" :active-index="activeIndex"></Header>
      </el-header>
      <el-container class="main-container">
        <Aside :active-index="activeIndex" ></Aside>
        <el-container>
          <el-main>
            <div class="main-content">
              <router-view />

            </div>
          </el-main>
        </el-container>
      </el-container>
    </el-container>
  </div>
</template>

<style scoped lang="scss">
.el-header{
  --el-header-padding:0;
}
.main-container{
  height: calc(100vh - 60px );
  overflow: hidden;
  overflow-y: scroll;
  background: #ddd;
  .el-main{
    padding: 20px 15px;

    .main-content{
      background: #fff;
      width: 100%;
      height: 100%;
      overflow-y: scroll;
    }
  }
}
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