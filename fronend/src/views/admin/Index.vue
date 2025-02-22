<script setup lang="ts">
import {ArrowRight} from "@element-plus/icons-vue";
import {onMounted, reactive, ref} from "vue";
import {delAdmin, getAdminList} from "@/assets/api/admin";
import {ElMessageBox, ElNotification} from "element-plus";
import router from "@/router";
const loading = ref(true)
const  adminData = ref([])
const  page = ref({
  total:0,
  page:0,
  size:15,
  next:false
})

const StatusSelect = ref([{
  "label":"禁止",
  "value":0,
},{
  "label":"正常",
  "value":1,
}])

const  searchFrom = reactive({
  email:"",
  username:"",
  phone:"",
  status:"",
  page:0,
  avatar:""

})
onMounted(async  ()=>{
    try {
       adminList("",1)
    }catch (e) {
      
    }
});
const handleSizeChange = (val: number) => {
  adminList(searchFrom,val)
}
const  handleEdit=(index:number, row:object)=>{
    router.push({path:"/admin/update/"+row.id})
}

const handleDelete=  (index:number, row:object) =>{
  ElMessageBox.confirm("确定要删除数据",{  confirmButtonText: '确定', cancelButtonText: '取消',}).then( async ()=>{
    loading.value=true
    const res =  await  delAdmin(row.id)
    if (res.code ==0){
      ElNotification.success("删除成功,3s之后刷新页面")
      setTimeout(function (){
        adminList(searchFrom,1)
      },3000)
    }
  }).catch(()=>{});
}
const handleCurrentChange = (val: number) => {
  adminList(searchFrom,val)
}
const onSearch =  ()=>{
  adminList(searchFrom,1)
}
const  onSave=()=>{
   router.push({"path":"/admin/save"})
}
const  adminList =async (data:any, val:number)=>{
  try {
    searchFrom.page=val
    const  res =await getAdminList(data)
    adminData.value=res.data 
    page.value = res.page
    loading.value=false
  }catch (e) {
    loading.value=false
  }
 
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
          <el-select v-model="searchFrom.status">
            <el-option
                v-for="item in StatusSelect"
                :key="item.value"
                :label="item.label"
                :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSearch">查询</el-button>
          <el-button type="primary" @click="onSave">新增</el-button>
        </el-form-item>
      </el-form>

    <el-table
        v-loading="loading"
        element-loading-text="Loading..."
        element-loading-svg-view-box="-10, -10, 50, 50"
        element-loading-background="rgba(122, 122, 122, 0.8)"
        :data="adminData"
        style="width: 100%"
        border="true"
    >
      <el-table-column prop="username" label="用户名"  />
      <el-table-column prop="phone" label="手机号"  />
      <el-table-column prop="email" label="邮箱" />
      <el-table-column prop="status" label="状态" >
        <template #default="scope">
           <el-link :type="scope.row.status == '正常'?'success':'danger'">{{scope.row.status}}</el-link>
        </template>
      </el-table-column>
      <el-table-column prop="img_url" label="头像">
        <template  #default="scope">
          <img :src="scope.row.img_url" alt="头像" style="max-height: 30px">
        </template>
      </el-table-column>
      <el-table-column label="操作" width="180">
        <template  #default="scope">
          <el-button size="small"  type="primary" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
          <el-button  size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
        </template>

      </el-table-column>
    </el-table>
      <el-pagination background layout="prev, pager, next" :total="page.total"
                     :page = page.page
                     :default-page-size = page.size
                     :hide-on-single-page=page.next
                     @size-change="handleSizeChange"
                     @current-change="handleCurrentChange"
      />
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
.el-pagination{
  margin-top: 20px;
}

#admin-index{
  .table-box{
    padding: 20px;
  }
}
.dialog-footer button:first-child {
  margin-right: 10px;
}
</style>