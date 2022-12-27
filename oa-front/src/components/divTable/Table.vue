<template>
    <el-table :data="data" :row-style="columnObj.rowStyle"
        :highlight-current-row="columnObj.highlightCurrentRow || false">

        <el-table-column v-for="(header, hi) in columnObj.headers" :key="hi" :prop="header.prop" :label="header.label"
            :min-width="header.width" :align="header.align || 'center'" :index="hi" show-overflow-tooltip>

            <template #default="scope">

                <!-- 是否 -->
                <span v-if="header.type === 'boolean'">
                    <div v-if="scope.row[header.prop] == true">是</div>
                    <div v-if="scope.row[header.prop] == false">否</div>
                </span>

                <!-- 按钮 -->
                <!-- <span v-if="header.type === 'operation'" v-for="(operation, oi) in header.operations" :key="oi"> -->
                <!-- <el-button v-if="operation.isShow()" :icon="operation.icon" :type="operation.type"
                        @click="operation.buttonClick(scope)" :style="{color:operation.color}" size="small">
                        {{operation.label}}
                    </el-button> -->
                <span v-if="header.type === 'operation'" v-for="operation in header.operations" :key="operation.id">
                    <el-button v-if="operation.isShow(scope.$index, scope.row)" :type="operation.type" plain
                        @click="operation.onClick(scope.$index, scope.row)" class="operation" :size="operation.size">
                        {{ operation.label }}
                    </el-button>
                </span>
            </template>

        </el-table-column>
    </el-table>
</template>
  
<script setup>
const props = defineProps({
    columnObj: {
        type: Object,
        default: {
            headers: []
        }
    },
    data: {
        type: Array,
        default: []
    }
})
</script>

<style>
.operation {
    margin: 0 3px;
}
</style>