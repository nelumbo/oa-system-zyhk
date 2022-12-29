<template>
    <div>
        <el-divider content-position="left">
            <h2>业务员</h2>
        </el-divider>
        <predegisnManage />
    </div>
    <div style="margin-top: 30px;">
        <el-divider content-position="left">
            <h2>技术负责人</h2>
        </el-divider>
        <predesignTaskManage />
    </div>
</template>

<script setup>
import predegisnManage from './PredegisnManage.vue'
import predesignTaskManage from './PredesignTaskManage.vue'
</script>