<script setup lang="ts">
import {inject, reactive, ref, watch} from 'vue';
import { ArrowRight } from '@element-plus/icons-vue';
import {ElForm, ElMessage} from 'element-plus';
import { ChangePassword } from "@/assets/api/admin";
import router from "@/router";
const  rules = reactive({
  old_password:[{ required: true, message: '请输入旧密码', trigger: 'blur' }],
  password:[{ required: true, message: '请输入密码', trigger: 'blur' }],
  confirm_password:[{required: true, message: '请输入确认密码', trigger: 'blur'}]
})

const formRef = ref<InstanceType<typeof ElForm> | null>(null);
const adminInfo = reactive({
   "old_password":"",
   "password":"",
   "confirm_password":"",
})
const onSubmit = async () => {
  try {
    if (!formRef.value ||(!(await formRef.value.validate()))) return;
    const res = await  ChangePassword(adminInfo)
    if (res.code ===0){
      localStorage.clear()
      ElMessage.success("修改密码成功,3s后请重新登录")
      setTimeout(function (){
         router.replace({"path":"/"})
      },3000)
    }else{
      ElMessage.error("修改密码失败")
    }
  }catch (e) {
    if(e.message)ElMessage.error("网络错误")
  }
};
</script>

<template>
  <div id="admin-info">
    <el-breadcrumb :separator-icon="ArrowRight">
      <el-breadcrumb-item :to="{ path: '/admin' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>修改密码</el-breadcrumb-item>
    </el-breadcrumb>
    <el-form :model="adminInfo" label-width="120px" class="info-form" size="large" ref="formRef" :rules="rules">
      <el-form-item label="旧密码" prop="old_password" >
        <el-input v-model="adminInfo.old_password"  show-password/>
      </el-form-item>
      <el-form-item label="新密码" prop="password">
        <el-input v-model="adminInfo.password"  show-password/>
      </el-form-item>
      <el-form-item label="确认密码" prop="confirm_password">
        <el-input v-model="adminInfo.confirm_password" show-password />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">确定</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<style lang="scss">
.avatar-uploader .avatar {
  width: 178px;
  height: 178px;
  display: block;
}
.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
}
#admin-info {
  .el-breadcrumb {
    font-size: 16px;
    margin-bottom: 20px;
    border-bottom: 1px solid #f4f4f4;
    padding: 20px;
    background: #faf7f7;
  }
  .info-form {
    padding: 20px 0;
    max-width: 600px;
  }
}
</style>