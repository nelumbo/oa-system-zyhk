import { ElMessage } from 'element-plus'

/**
 * 消息提示
 * @param msg 提示信息
 * @param type 消息类型
 */
export const message = (msg, type) => {
    ElMessage({
        showClose: true,
        message: msg,
        type: type
    });
}

/**
 * 消息提示
 * @param name 消息主体名称
 * @param msg 提示信息
 * @param type 消息类型
 */
export const messageForCRUD = (name, msg, type) => {
    msg = "【" + name + "】" + msg
    ElMessage({
        showClose: true,
        message: msg,
        type: type
    });
}