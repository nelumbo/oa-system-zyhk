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


                <!-- 特殊 -->
                <!-- 合同未回款额计算 -->
                <span v-if="header.type === 'contractNotPayment'">
                    {{ scope.row.isPreDeposit ? "-" : scope.row.totalAmount - scope.row.paymentTotalAmount }}
                </span>

                <!-- 合同预存款展示 -->
                <span v-if="header.type === 'contractPreDeposit'">
                    {{ scope.row.isPreDeposit ? scope.row[header.prop] : "-" }}
                </span>

                <span v-if="header.type === 'contractType'">
                    {{ contractStatusToText(scope.row) }}
                </span>

                <span v-if="header.type === 'employees'">
                    <div v-if="scope.row.technicianMan.name">
                        技术：{{ scope.row.technicianMan.name }}
                    </div>
                    <div v-if="scope.row.purchaseMan.name">
                        采购：{{ scope.row.purchaseMan.name }}
                    </div>
                    <div v-if="scope.row.inventoryMan.name">
                        仓库：{{ scope.row.inventoryMan.name }}
                    </div>
                    <div v-if="scope.row.shipmentMan.name">
                        物流：{{ scope.row.shipmentMan.name }}
                    </div>
                </span>


                <span v-if="header.type === 'taskID'">
                    {{ scope.row.task.id == 0 ? "" : scope.row.task.id }}
                </span>


                <span v-if="header.type === 'taskStartDate'">
                    <div v-if="scope.row.type == 3">技术：{{ scope.row.technicianStartDate }}</div>
                    <div v-if="scope.row.type > 1 && scope.row.status > 1">
                        采购：{{ scope.row.purchaseStartDate }}
                    </div>
                    <div v-if="scope.row.status > 2">仓库：{{ scope.row.inventoryStartDate }}</div>
                    <div v-if="scope.row.type == 3 && scope.row.status > 3">装配：{{ scope.row.assemblyStartDate }}</div>
                    <div v-if="scope.row.status > 4">物流：{{ scope.row.shipmentStartDate }}</div>
                </span>

                <span v-if="header.type === 'taskDays'">
                    <div v-if="scope.row.type == 3">技术：{{ scope.row.technicianDays }}天</div>
                    <div v-if="scope.row.type > 1"> 采购：{{ scope.row.purchaseDays }}天</div>
                </span>

                <span v-if="header.type === 'taskFinalDate'">
                    <div v-if="scope.row.type == 3 && scope.row.status > 1">技术：{{ scope.row.technicianFinalDate }}</div>
                    <div v-if="scope.row.type > 1 && scope.row.status > 2"> 采购：{{ scope.row.purchaseFinalDate }}</div>
                    <div v-if="scope.row.status > 3">仓库：{{ scope.row.inventoryFinalDate }}</div>
                    <div v-if="scope.row.type == 3 && scope.row.status > 4">装配：{{ scope.row.assemblyFinalDate }}</div>
                    <div v-if="scope.row.status > 5">物流：{{ scope.row.shipmentFinalDate }}</div>
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
import { contractStatusItems, productionStatusItems, collectionStatusItems } from '@/utils/magic'
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

function contractStatusToText(row) {
    if (row.status == 2) {
        return (
            statusToText(productionStatusItems, row.productionStatus) +
            "," +
            statusToText(collectionStatusItems, row.collectionStatus)
        );
    } else {
        return statusToText(contractStatusItems, row.status);
    }
}
</script>

<style>
.operation {
    margin: 0 3px;
}
</style>