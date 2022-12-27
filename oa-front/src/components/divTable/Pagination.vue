<template>
    <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :hide-on-single-page="false"
        :total="pageData.total" :page-size="pageData.pageSize" :current-page="pageData.pageNo" :page-sizes="[10, 20, 30]"
        layout="total,prev,pager,next, jumper,sizes" :style="{justifyContent: 'right'}">
    </el-pagination>
</template>

<script setup>
const props = defineProps({
    pageData: {
        type: Object,
        default: {
            total: 0,
            pageSize: 10,
            pageNo: 1,
        }
    },
    handleSizeChange: {
        type: Function,
    },
    handleCurrentChange: {
        type: Function
    }
})
</script>