<script setup lang="ts">
import {inject, reactive, ref, watch} from 'vue';
import { ArrowRight, Plus } from '@element-plus/icons-vue';
import type { UploadProps } from 'element-plus';
import {ElForm, ElMessage} from 'element-plus';
import {ChangeInfo} from "@/assets/api/admin";
const  rules = reactive({
  username:[{ required: true, message: '请输入用户名', trigger: 'blur' }],
  email:[{ required: true, message: '请输入邮箱', trigger: 'blur' }],
  phone:[{required: true, message: '请输入手机号', trigger: 'blur'}]
})

const adminInfo = inject<Record<string, any>>("adminInfo");
const formRef = ref<InstanceType<typeof ElForm> | null>(null);
const localAdminInfo = reactive({...(adminInfo || {}) });
const imageUrl = ref(localAdminInfo.img_url);


watch(
    ()=>adminInfo,
    (newVal) => {
      Object.assign(localAdminInfo, newVal);
      imageUrl.value =localAdminInfo.img_url
    },
    { deep: true }
)
const handleAvatarSuccess: UploadProps['onSuccess'] = (response) => {
   if (response.code===0){
     imageUrl.value= response.data.url;
     localAdminInfo.avatar = response.data.path;
   }else{
     ElMessage.error(response.data.msg)
   }
};

const beforeAvatarUpload: UploadProps['beforeUpload'] = (rawFile) => {
  if (rawFile.type !== 'image/jpeg') {
    ElMessage.error('Avatar picture must be JPG format!');
    return false;
  } else if (rawFile.size / 1024 / 1024 > 2) {
    ElMessage.error('Avatar picture size can not exceed 2MB!');
    return false;
  }
  return true;
};

const onSubmit = async () => {
  try {
    if (!formRef.value) return
    const isValid = await formRef.value.validate();
    if (!isValid) return;
    const res = await ChangeInfo(localAdminInfo)
    if (res.code === 0){
      ElMessage.success(res.msg)
    }else{
      ElMessage.error(res.msg)
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
      <el-breadcrumb-item>我的资料</el-breadcrumb-item>
    </el-breadcrumb>
    <el-form :model="localAdminInfo" label-width="120px" class="info-form" size="large" ref="formRef" :rules="rules">
      <el-form-item label="用户名" prop="username">
        <el-input v-model="localAdminInfo.username" />
      </el-form-item>
      <el-form-item label="邮箱" prop="email">
        <el-input v-model="localAdminInfo.email" />
      </el-form-item>
      <el-form-item label="手机号" prop="phone">
        <el-input v-model="localAdminInfo.phone" />
      </el-form-item>
      <el-form-item label="头像">
        <el-upload
            class="avatar-uploader"
            action="http://localhost:8078/uploads/file"
            :show-file-list="false"
            :on-success="handleAvatarSuccess"
            :before-upload="beforeAvatarUpload"
        >
          <img v-if="imageUrl" :src="imageUrl" class="avatar" />
          <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
        </el-upload>
      </el-form-item>
      <el-form-item label="状态">
        <el-select v-model="localAdminInfo.status" disabled>
          <el-option value="禁止">禁止</el-option>
          <el-option value="正常">正常</el-option>
        </el-select>
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

  .info-form {
    padding: 20px 0;
    max-width: 600px;
  }
}
</style>