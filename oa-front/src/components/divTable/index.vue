<template>
    <elTable :columnObj="columnObj" :data="tableData"></elTable>
    <div style="margin-top: 5px;"></div>
    <elPagination :pageData="pageData" :handleSizeChange="handleSizeChange" :handleCurrentChange="handleCurrentChange">
    </elPagination>
</template>

<script setup>
import elTable from './Table.vue'
import elPagination from './Pagination.vue'

const props = defineProps({
    columnObj: {
        type: Object,
    },
    tableData: {
        type: Array,
        default: {
            total: 0,
            pageSize: 10,
            pageNo: 1,
        }
    },
    pageData: {
        type: Object,
    },
    handleSizeChange: {
        type: Function
    },
    handleCurrentChange: {
        type: Function
    }
})

</script>