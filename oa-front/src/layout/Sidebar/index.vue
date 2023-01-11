<template>
  <el-menu class="menu" active-text-color="#ffd04b" background-color="#545c64" :default-active="$route.path"
    text-color="#fff" :collapse="viewType().isMobbile" router>
    <el-sub-menu index="1">
      <template #title>
        <el-icon>
          <ElementPlus />
        </el-icon>
        <span class="title">{{ user().my.name }}</span>
      </template>
      <el-menu-item-group>
        <el-menu-item index="1-1">
          <template #title>
            <el-icon>
              <Setting />
            </el-icon>
            <span>设置</span>
          </template>
        </el-menu-item>
        <el-menu-item index="1-2" @click="exit">
          <template #title>
            <el-icon>
              <Close />
            </el-icon>
            <span>退出</span>
          </template>
        </el-menu-item>
      </el-menu-item-group>
    </el-sub-menu>
    <el-divider />
    <el-menu-item index="/home">
      <el-icon>
        <House />
      </el-icon>
      <span>首页</span>
    </el-menu-item>
    <el-menu-item index="/expense">
      <el-icon>
        <Money />
      </el-icon>
      <span>我的报销</span>
    </el-menu-item>
    <el-menu-item index="/contract">
      <el-icon>
        <Files />
      </el-icon>
      <span>合同管理</span>
    </el-menu-item>
    <el-menu-item index="/pre-research">
      <el-icon>
        <List />
      </el-icon>
      <span>预设计管理</span>
    </el-menu-item>
    <el-menu-item index="/customer">
      <el-icon>
        <User />
      </el-icon>
      <span>客户管理</span>
    </el-menu-item>
    <el-menu-item index="/product">
      <el-icon>
        <Box />
      </el-icon>
      <span>产品管理</span>
    </el-menu-item>
    <el-menu-item index="/purchasing">
      <el-icon>
        <ShoppingTrolley />
      </el-icon>
      <span>采购管理</span>
    </el-menu-item>
    <el-menu-item index="/finance">
      <el-icon>
        <Ticket />
      </el-icon>
      <span>财务管理</span>
    </el-menu-item>
    <el-menu-item index="/enterprise">
      <el-icon>
        <Avatar />
      </el-icon>
      <span>企业管理</span>
    </el-menu-item>
    <el-menu-item index="/log">
      <el-icon>
        <Checked />
      </el-icon>
      <span>日志系统</span>
    </el-menu-item>
  </el-menu>
</template>

<script setup>
import { viewType } from '@/pinia/modules/viewType'
import { user } from '@/pinia/modules/user'

function exit() {
  localStorage.removeItem("token");
  window.location.href = '/login'
}

</script>

<style>
.menu {
  border-right-width: 0
}

.title {
  font-size: 30px;
  font-weight: normal;
  margin: 20px 0 0;
}
</style>