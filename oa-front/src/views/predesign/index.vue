<template>
    <div v-if="user().my.pids.includes('21')">
        <el-divider content-position="left">
            <h2>业务员</h2>
        </el-divider>
        <predegisnManage />
    </div>
    <div style="margin-top: 30px;" v-if="user().my.pids.includes('23')">
        <el-divider content-position="left">
            <h2>技术负责人</h2>
        </el-divider>
        <predesignTaskManage />
    </div>
</template>

<script setup>
import { user } from '@/pinia/modules/user'
import predegisnManage from './PredegisnManage.vue'
import predesignTaskManage from './PredesignTaskManage.vue'
</script>