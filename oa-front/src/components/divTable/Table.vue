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

                <span v-if="header.type === 'transform'">
                    {{ statusToText(header.items, scope.row[header.prop]) }}
                </span>

                <span v-if="header.type === 'textarea'"
                    @click="textarea.openTextareaDialog(header.label, scope.row[header.prop])">
                    {{ scope.row[header.prop] }}
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
    <el-dialog v-model="textarea.dialogVisible" :title="textarea.model.title" width="50%" :show-close="false">
        <el-form :model="textarea.model" label-width="80px">
            <el-input v-model="textarea.model.text" type="textarea" autosize readonly />
        </el-form>
    </el-dialog>
</template>
  
<script setup>
import { reactive } from 'vue'
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

const textarea = reactive({
    dialogVisible: false,
    model: {
        title: "",
        text: ""
    },
    openTextareaDialog: (title, text) => {
        textarea.model.title = title
        textarea.model.text = text
        textarea.dialogVisible = true
    }
})

function statusToText(items, status) {
    var temp = "";
    items.some((item) => {
        if (item.value == status) {
            temp = item.label;
            return;
        }
    });
    return temp;
}
</script>

<style>
.operation {
    margin: 0 3px;
}
</style>