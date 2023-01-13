<template>
    <div v-if="user().my.pids.includes('1') || user().my.pids.includes('2')">
        <el-divider content-position="left">
            <h2>排行榜</h2>
        </el-divider>
        <ranking />
    </div>
    <div style="margin-top: 30px;" v-if="user().my.pids.includes('3')">
        <el-divider content-position="left">
            <h2>我的合同</h2>
        </el-divider>
        <contractManage />
    </div>
    <div style="margin-top: 30px" v-if="
    user().my.pids.includes('4') ||
    user().my.pids.includes('5') ||
    user().my.pids.includes('6') ||
    user().my.pids.includes('7')">
        <el-divider content-position="left">
            <h2>我的任务</h2>
        </el-divider>
        <taskManage />
    </div>
    <div style="margin-top: 30px;" v-if="
    user().my.pids.includes('8') ||
    user().my.pids.includes('9') ||
    user().my.pids.includes('10')">
        <el-divider content-position="left">
            <h2>我的回款追踪任务</h2>
        </el-divider>
        <paymentManage />
    </div>
    <div style="margin-top: 30px;" v-if="user().my.pids.includes('11')">
        <el-divider content-position="left">
            <h2>我发起的预设计</h2>
        </el-divider>
        <predesignManage />
    </div>
    <div style="margin-top: 30px;" v-if="user().my.pids.includes('12')">
        <el-divider content-position="left">
            <h2>我的预设计任务</h2>
        </el-divider>
        <predesignTaskManage />
    </div>
    <div style="margin-top: 30px;" v-if="user().my.pids.includes('3')">
        <el-divider content-position="left">
            <h2>我的投标保证金</h2>
        </el-divider>
        <bidBondManage />
    </div>
</template>

<script setup>
import { user } from '@/pinia/modules/user'
import ranking from './Ranking.vue';
import contractManage from './ContractManage.vue'
import taskManage from './TaskManage.vue'
import paymentManage from './PaymentManage.vue'
import predesignManage from './PredesignManage.vue'
import predesignTaskManage from './PredesignTaskManage.vue'
import bidBondManage from './BidBondManage.vue'
</script>