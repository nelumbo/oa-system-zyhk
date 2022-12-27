<template>
    <div>
        <el-divider content-position="left">
            <h2>业务员</h2>
        </el-divider>
        <preResearchManage />
    </div>
    <div style="margin-top: 30px;">
        <el-divider content-position="left">
            <h2>技术负责人</h2>
        </el-divider>
        <preResearchTaskManage />
    </div>
</template>

<script setup>
import preResearchManage from './PreResearchManage.vue'
import preResearchTaskManage from './PreResearchTaskManage.vue'
</script>