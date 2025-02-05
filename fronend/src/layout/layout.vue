<script setup lang="ts">
import {onMounted, ref, watch} from 'vue'
import { useRouter } from 'vue-router';
import Aside from "@/components/aside.vue";
import Header from "@/components/header.vue";
import Footer from "@/components/footer.vue";
const chooseHeader = window.location.pathname;
// 创建一个响应式引用并初始化为当前页面的路径
const activeIndex = ref<string>(chooseHeader);
const router = useRouter();
// 监听路由变化并更新 activeIndex
onMounted(() => {
  watch(
      () => router.currentRoute.value.path,
      (newPath) => {
        activeIndex.value = newPath;
      }
  );
});
</script>
<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <Header :active-index="activeIndex"></Header>
      </el-header>
      <el-container class="main-container">
        <Aside :active-index="activeIndex"></Aside>
        <el-container>
          <el-main>
            <div class="main-content">

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
    }
  }
}


</style>