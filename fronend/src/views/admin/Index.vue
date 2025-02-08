<script setup lang="ts">
import {ArrowRight} from "@element-plus/icons-vue";
import {onMounted, reactive, ref} from "vue";
import {getAdminList} from "@/assets/api/admin";

const  adminData = ref([])


const  searchFrom = reactive({
  email:"",
  username:"",
  phone:"",
  status:"",
  page:"",
  avatar:""

})

onMounted(async  ()=>{
    try {
        const  res =await getAdminList([])
        adminData.value=res.data
    }catch (e) {
      
    }
});

const onSearch = async  ()=>{
  const  res =await getAdminList(searchFrom)
  adminData.value=res.data
}

</script>

<template>
  <div id="admin-index">
   <el-breadcrumb :separator-icon="ArrowRight">
      <el-breadcrumb-item :to="{ path: '/admin' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>管理员管理</el-breadcrumb-item>
      <el-breadcrumb-item>管理员列表</el-breadcrumb-item>
    </el-breadcrumb>
    <div class="table-box">

      <el-form :inline="true" :model="searchFrom" class="demo-form-inline">
        <el-form-item label="用户名:">
          <el-input v-model="searchFrom.username" placeholder="用户名" clearable />
        </el-form-item>
        <el-form-item label="手机号:">
          <el-input v-model="searchFrom.phone" placeholder="手机号" clearable />
        </el-form-item>
        <el-form-item label="邮箱:">
          <el-input v-model="searchFrom.email" placeholder="邮箱" clearable />
        </el-form-item>
        <el-form-item label="状态:">
          <el-input v-model="searchFrom.status" placeholder="状态" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSearch">查询</el-button>
        </el-form-item>
      </el-form>
    <el-table
        :data="adminData"
        style="width: 100%"
        border="true"
    >
      <el-table-column prop="username" label="用户名"  />
      <el-table-column prop="phone" label="手机号"  />
      <el-table-column prop="email" label="邮箱" />
      <el-table-column prop="status" label="状态" />
      <el-table-column prop="img_url" label="头像">
        <template  #default="scope">
          <img :src="scope.row.img_url" alt="头像" style="max-height: 30px">
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template  #default="scope">
          <a href="">编辑</a>
          <a href="">删除</a>
        </template>
      </el-table-column>
    </el-table>
    </div>
  </div>
</template>

<style scoped lang="scss">
.demo-form-inline {
  .el-input {
    --el-input-width: 250px;
  }
    .el-select {
      --el-select-width: 250px;
    }
}

#admin-index{
  .table-box{
    padding: 20px;
  }
}
</style>