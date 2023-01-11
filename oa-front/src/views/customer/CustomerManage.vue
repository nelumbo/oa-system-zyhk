<template>
    <el-row :gutter="20">
        <el-col :span="4" :offset="3">
            <el-select v-model="base.model.customerCompany.regionID" placeholder="区域" clearable style="width: 100%;">
                <el-option v-for="item in base.regions" :key="item.id" :label="item.name" :value="item.id" />
            </el-select>
        </el-col>
        <el-col :span="4">
            <el-input v-model="base.model.customerCompany.name" placeholder="公司" clearable maxlength="25" />
        </el-col>
        <el-col :span="4">
            <el-input v-model="base.model.name" placeholder="姓名" clearable maxlength="25" />
        </el-col>
        <el-col :span="4">
            <el-input v-model="base.model.researchGroup" placeholder="课题组" clearable maxlength="25" />
        </el-col>
        <el-col :span="1">
            <el-button type="primary" @click="base.query">查询</el-button>
        </el-col>
    </el-row>
    <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
        :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />

    <el-dialog v-model="edit.dialogVisible" title="客户编辑" width="50%" :show-close="false">
        <el-form :model="edit.model" label-width="100px" :rules="rules" ref="editForm">
            <el-form-item label="公司">
                <el-input v-model="edit.model.customerCompanyName" disabled />
            </el-form-item>
            <el-form-item label="姓名" prop="name">
                <el-input v-model.trim="edit.model.name" maxlength="50" />
            </el-form-item>
            <el-form-item label="课题组">
                <el-input v-model.trim="edit.model.researchGroup" maxlength="100" />
            </el-form-item>
            <el-form-item label="电话">
                <el-input v-model.trim="edit.model.phone" maxlength="50" />
            </el-form-item>
            <el-form-item label="微信号">
                <el-input v-model.trim="edit.model.wechatID" maxlength="50" />
            </el-form-item>
            <el-form-item label="电子邮箱">
                <el-input v-model.trim="edit.model.email" maxlength="50" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="primary" @click="edit.submit" :disabled="edit.submitDisabled">提交</el-button>
                </div>
            </span>
        </template>
    </el-dialog>

    <el-dialog v-model="del.dialogVisible" title="删除" width="50%" :show-close="false">
        <h1>是否确定删除【{{ del.model.name }}】？</h1>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="danger" @click="del.submit" :disabled="del.submitDisabled">确定</el-button>
                </div>
            </span>
        </template>
    </el-dialog>
</template>

<script setup>
import { ref, reactive, onBeforeMount } from 'vue'
import { delCustomer, editCustomer, queryCustomers } from "@/api/customer"
import { queryAllRegion } from '@/api/region'
import { message, messageForCRUD } from '@/components/divMessage/index'

import divTable from '@/components/divTable/index.vue'

const editForm = ref(null)
const rules = reactive({
    name: [
        { required: true, message: '请输入客户名称', trigger: 'blur' },
    ],
})

const base = reactive({
    regions: [],
    model: {
        name: "",
        researchGroup: "",
        customerCompany: {
            regionID: null,
            name: "",
        },
    },
    column: {
        headers: [
            {
                prop: "customerCompany.region.name",
                label: "区域",
            },
            {
                prop: "customerCompany.name",
                label: "公司",
            },
            {
                prop: "name",
                label: "姓名",
            },
            {
                prop: "researchGroup",
                label: "课题组",
            },
            {
                prop: "phone",
                label: "联系电话",
            },
            {
                prop: "wechatID",
                label: "微信号",
            },
            {
                prop: "email",
                label: "电子邮箱",
            },
            {
                type: "operation",
                label: "操作",
                operations: [
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "编辑",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openEditDialog(index, row)
                    },
                    // {
                    //     isShow: (index, row) => {
                    //         return true
                    //     },
                    //     label: "删除",
                    //     type: "danger",
                    //     align: "center",
                    //     sortable: false,
                    //     onClick: (index, row) => base.openDelDialog(index, row)
                    // }
                ]
            },
        ],
    },
    tableData: [],
    pageData: {
        total: 0,
        pageSize: 10,
        pageNo: 1
    },
    query: () => {
        queryCustomers(base.model, base.pageData).then((res) => {
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
        base.pageData.pageSize = e
        base.pageData.pageNo = 1
        base.query()
    },
    handleCurrentChange: (e) => {
        base.pageData.pageNo = e
        base.query()
    },
    openEditDialog: (index, row) => {
        edit.model.id = row.id
        edit.model.customerCompanyID = row.customerCompanyID
        edit.model.customerCompanyName = row.customerCompany.name
        edit.model.name = row.name
        edit.model.regionID = row.regionID
        edit.model.researchGroup = row.researchGroup
        edit.model.phone = row.phone
        edit.model.wechatID = row.wechatID
        edit.model.email = row.email
        edit.dialogVisible = true
    },
    openDelDialog: (index, row) => {
        del.model.id = row.id
        del.model.name = row.name
        del.dialogVisible = true
    }
})

const edit = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: null,
        customerCompanyID: null,
        customerCompanyName: "",
        name: "",
        researchGroup: "",
        phone: "",
        wechatID: "",
        email: "",
    },
    submit: () => {
        editForm.value.validate((valid) => {
            if (valid) {
                edit.submitDisabled = true
                editCustomer(edit.model).then((res) => {
                    if (res.status == 1) {
                        messageForCRUD(edit.model.name, "编辑成功", "success")
                        base.query()
                    } else {
                        messageForCRUD(edit.model.name, "编辑失败", "error")
                    }
                    edit.dialogVisible = false
                    edit.model = {
                        id: null,
                        customerCompanyID: null,
                        customerCompanyName: "",
                        name: "",
                        researchGroup: "",
                        phone: "",
                        wechatID: "",
                        email: "",
                    }
                    edit.submitDisabled = false
                })
            } else {
                return false;
            }
        })
    }
})

const del = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: 0,
        name: "",
    },
    submit: () => {
        del.submitDisabled = true
        delCustomer(del.model.id).then((res) => {
            if (res.status == 1) {
                messageForCRUD(del.model.name, "删除成功", "success")
                base.query()
            } else {
                messageForCRUD(del.model.name, "删除失败，该客户录入过合同。", "error")
            }
            del.dialogVisible = false
            del.model = {
                id: 0,
                name: "",
            }
            del.submitDisabled = false
        })
    }
})

onBeforeMount(() => {
    queryAllRegion().then((res) => {
        if (res.status == 1) {
            base.regions = res.data
        }
    })
    base.query()
})

defineExpose({
    base
})
</script>