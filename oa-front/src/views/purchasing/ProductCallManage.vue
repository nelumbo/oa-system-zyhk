<template>
    <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
        :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />

    <el-dialog v-model="add.dialogVisible" title="添加" width="50%" :show-close="false">
        <el-form :model="add.model" label-width="100px" :rules="rules" ref="addForm">
            <el-form-item label="类型" prop="type">
                <el-select v-model="add.model.type" placeholder="请选择你的采购类型">
                    <el-option v-for="purchasing in PurchasingTypeSelectItems" :label="purchasing.label"
                        :value="purchasing.value" />
                </el-select>
            </el-form-item>
            <el-form-item label="数量" prop="realNumber">
                <el-input-number v-model="add.model.realNumber" :controls="false" :min="1" :max="99999" />
            </el-form-item>
            <el-form-item label="单价" prop="price">
                <el-input-number v-model="add.model.price" :controls="false" :min="0" :max="99999" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="primary" @click="add.submit" :disabled="add.submitDisabled">提交</el-button>
                </div>
            </span>
        </template>
    </el-dialog>
</template>

<script setup>
import { ref, reactive, onBeforeMount } from 'vue'
import { message } from '@/components/divMessage/index'
import { PurchasingTypeSelectItems } from '@/utils/magic'
import { addPurchasing } from "@/api/purchasing"
import { queryProductCalls } from "@/api/product_call"
import { reg_number, reg_money } from '@/utils/regex'

import divTable from '@/components/divTable/index.vue'


const addForm = ref(null)
const rules = reactive({
    type: [
        { required: true, message: '请选择', trigger: 'blur' },
    ],
    realNumber: [
        { required: true, pattern: reg_number, message: '请输入大于零的有效数字', trigger: 'blur' },
    ],
    price: [
        { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
    ],
})

const base = reactive({
    column: {
        headers: [
            {
                prop: "createDate",
                label: "产品",
            },
            {
                prop: "product.name",
                label: "产品",
            },
            {
                prop: "product.number",
                label: "库存数量",
            },
            {
                prop: "product.callNumber",
                label: "报警数量",
            },
            {
                type: "operation",
                label: "操作",
                operations: [
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "采购",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openAddDialog(index, row)
                    },
                ]
            },
        ]
    },
    tableData: [],
    pageData: {
        total: 0,
        pageSize: 10,
        pageNo: 1
    },
    query: () => {
        queryProductCalls(null, base.pageData).then((res) => {
            if (res.status == 1) {
                base.tableData = res.data.data
                base.pageData.total = res.data.total
                base.pageData.pageSize = res.data.pageSize
                base.pageData.pageNo = res.data.pageNo
            } else {
                message("查询失败", "error")
            }
        })
    },
    handleSizeChange: (e) => {
        base.pageData.pageSize = ae
        base.pageData.pageNo = 1
        base.query()
    },
    handleCurrentChange: (e) => {
        base.pageData.pageNo = e
        base.query()
    },
    openAddDialog(index, row) {
        add.model.productID = row.id
        add.dialogVisible = true
    },
})


const add = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        productID: null,
        type: null,
        price: 0,
        realNumber: 0,
    },
    submit: () => {
        addForm.value.validate((valid) => {
            if (valid) {
                add.submitDisabled = true
                addPurchasing(add.model).then((res) => {
                    if (res.status == 1) {
                        message("录入成功", "success")
                        base.query()
                    } else {
                        message("录入失败", "error")
                    }
                    add.dialogVisible = false
                    add.model = {
                        productID: null,
                        type: null,
                        price: 0,
                        realNumber: 0,
                    }
                    add.submitDisabled = false
                })
            } else {
                return false;
            }
        });
    },
})

onBeforeMount(() => {
    base.query()
})
</script>