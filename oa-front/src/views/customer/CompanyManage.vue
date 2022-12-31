<template>
    <div>
        <el-row :gutter="20">
            <el-col :span="6" :offset="5">
                <el-select v-model="base.model.regionID" placeholder="区域" clearable style="width: 100%;">
                    <el-option v-for="item in base.regions" :key="item.id" :label="item.name" :value="item.id" />
                </el-select>
            </el-col>
            <el-col :span="6">
                <el-input v-model="base.model.name" placeholder="名称" clearable maxlength="25" />
            </el-col>
            <el-col :span="1">
                <el-button type="primary" @click="base.query">查询</el-button>
            </el-col>
            <el-col :span="1">
                <el-button type="success" @click="base.openAddDialog">添加</el-button>
            </el-col>
        </el-row>
        <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
            :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />

        <el-dialog v-model="add.dialogVisible" title="客户公司添加" width="50%" :show-close="false">
            <el-form :model="add.model" label-width="60px" :rules="rules" ref="addForm">
                <el-form-item label="区域" prop="regionID">
                    <el-select v-model="add.model.regionID" clearable>
                        <el-option v-for="region in base.regions" :key="region.id" :label="region.name"
                            :value="region.id" />
                    </el-select>
                </el-form-item>
                <el-form-item label="名称" prop="name">
                    <el-input v-model.trim="add.model.name" maxlength="50" />
                </el-form-item>
                <el-form-item label="地址">
                    <el-input v-model.trim="add.model.address" maxlength="100" />
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

        <el-dialog v-model="add2.dialogVisible" title="客户添加" width="50%" :show-close="false">
            <el-form :model="add2.model" label-width="100px" :rules="rules2" ref="add2Form">
                <el-form-item label="公司名称">
                    <el-input v-model="add2.model.customerCompanyName" disabled />
                </el-form-item>
                <el-form-item label="姓名" prop="name">
                    <el-input v-model.trim="add2.model.name" maxlength="50" />
                </el-form-item>
                <el-form-item label="课题组">
                    <el-input v-model.trim="add2.model.researchGroup" maxlength="100" />
                </el-form-item>
                <el-form-item label="电话">
                    <el-input v-model.trim="add2.model.phone" maxlength="50" />
                </el-form-item>
                <el-form-item label="微信号">
                    <el-input v-model.trim="add2.model.wechatID" maxlength="50" />
                </el-form-item>
                <el-form-item label="电子邮箱">
                    <el-input v-model.trim="add2.model.email" maxlength="50" />
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="add2.submit" :disabled="add2.submitDisabled">提交</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="edit.dialogVisible" title="客户公司编辑" width="50%" :show-close="false">
            <el-form :model="edit.model" label-width="100px" :rules="rules" ref="editForm">
                <el-form-item label="区域" prop="regionID">
                    <el-select v-model="edit.model.regionID" clearable>
                        <el-option v-for="region in base.regions" :key="region.id" :label="region.name"
                            :value="region.id" />
                    </el-select>
                </el-form-item>
                <el-form-item label="名称" prop="name">
                    <el-input v-model.trim="edit.model.name" maxlength="50" />
                </el-form-item>
                <el-form-item label="地址">
                    <el-input v-model.trim="edit.model.address" maxlength="100" />
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

        <el-dialog v-model="del.dialogVisible" title="客户公司删除" width="50%" :show-close="false">
            <h1>是否确定删除【{{ del.model.name }}】？</h1>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="danger" @click="del.submit" :disabled="del.submitDisabled">确定</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>
    </div>
</template>

<script setup>
import { ref, reactive, onBeforeMount } from 'vue'
import { addCustomerCompany, delCustomerCompany, editCustomerCompany, queryCustomerCompanys } from "@/api/customer_company"
import { addCustomer } from "@/api/customer"
import { queryAllRegion } from '@/api/region'
import { message, messageForCRUD } from '@/components/divMessage/index'

import divTable from '../../components/divTable/index.vue'

const props = defineProps({
    handleAddCustomer: {
        type: Function,
    },
})

const addForm = ref(null)
const add2Form = ref(null)
const editForm = ref(null)

const rules = reactive({
    regionID: [
        { required: true, message: '请选择区域', trigger: 'blur' },
    ],
    name: [
        { required: true, message: '请输入客户公司名称', trigger: 'blur' },
    ],
})
const rules2 = reactive({
    name: [
        { required: true, message: '请输入客户名称', trigger: 'blur' },
    ],
})

const base = reactive({
    regions: [],
    model: {
        regionID: null,
        name: "",
    },
    column: {
        headers: [
            {
                prop: "region.name",
                label: "区域",
            },
            {
                prop: "name",
                label: "名称",
            },
            {
                prop: "address",
                label: "地址",
            },
            {
                type: "operation",
                label: "操作",
                operations: [
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "添加客户",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openAdd2Dialog(index, row)
                    },
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
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "删除",
                        type: "danger",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openDelDialog(index, row)
                    }
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
        queryCustomerCompanys(base.model, base.pageData).then((res) => {
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
    openAddDialog: () => {
        add.dialogVisible = true
    },
    openAdd2Dialog: (index, row) => {
        add2.model.customerCompanyID = row.id
        add2.model.customerCompanyName = row.name
        add2.dialogVisible = true
    },
    openEditDialog: (index, row) => {
        edit.model.id = row.id
        edit.model.regionID = row.regionID
        edit.model.name = row.name
        edit.model.address = row.address
        edit.dialogVisible = true
    },
    openDelDialog: (index, row) => {
        del.model.id = row.id
        del.model.name = row.name
        del.dialogVisible = true
    }
})

const add = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        regionID: null,
        name: "",
        address: "",
    },
    submit: () => {
        addForm.value.validate((valid) => {
            if (valid) {
                add.submitDisabled = true
                addCustomerCompany(add.model).then((res) => {
                    if (res.status == 1) {
                        messageForCRUD(add.model.name, "添加成功", "success")
                        base.query()
                    } else {
                        messageForCRUD(add.model.name, "添加失败", "error")
                    }
                    add.dialogVisible = false
                    add.model = {
                        regionID: null,
                        name: "",
                        address: "",
                    }
                    add.submitDisabled = false
                })
            } else {
                return false;
            }
        });
    }
})

const add2 = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        customerCompanyID: null,
        customerCompanyName: null,
        name: "",
        researchGroup: "",
        phone: "",
        wechatID: "",
        email: "",
    },
    submit: () => {
        add2Form.value.validate((valid) => {
            if (valid) {
                add2.submitDisabled = true
                addCustomer(add2.model).then((res) => {
                    if (res.status == 1) {
                        messageForCRUD(add2.model.name, "添加成功", "success")
                        props.handleAddCustomer()
                        base.query()
                    } else {
                        messageForCRUD(add2.model.name, "添加失败", "error")
                    }
                    add2.dialogVisible = false
                    add2.model = {
                        customerCompanyID: null,
                        customerCompanyName: null,
                        name: "",
                        researchGroup: "",
                        phone: "",
                        wechatID: "",
                        email: "",
                    }
                    add2.submitDisabled = false
                })
            } else {
                return false;
            }
        });
    }
})

const edit = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: 0,
        regionID: null,
        name: "",
        address: "",
    },
    submit: () => {
        editForm.value.validate((valid) => {
            if (valid) {
                edit.submitDisabled = true
                editCustomerCompany(edit.model).then((res) => {
                    if (res.status == 1) {
                        messageForCRUD(edit.model.name, "编辑成功", "success")
                        base.query()
                    } else {
                        messageForCRUD(edit.model.name, "编辑失败", "error")
                    }
                    edit.dialogVisible = false
                    edit.model = {
                        id: 0,
                        regionID: null,
                        name: "",
                        address: "",
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
        delCustomerCompany(del.model.id).then((res) => {
            if (res.status == 1) {
                messageForCRUD(del.model.name, "删除成功", "success")
                base.query()
            } else {
                messageForCRUD(del.model.name, "删除失败，该公司已经添加了客户。", "error")
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
</script>