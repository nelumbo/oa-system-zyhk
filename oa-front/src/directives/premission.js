export default function (app) {
    app.directive('permission', {
        mounted(el, bindings) {
            // 获取到指令的值
            let bv = bindings.value;
            // 判断当前是否存在值
            if (!bv) return console.error("自定义指令值缺失");

            // 是否为管理者：这里设置固定值,
            let isSuperAdmin = false;
            // 权限集合
            let permiList = ["C", "R", "U", "D"];

            // 判断当前按钮值是否存在权限集合中并且当前用户是否为超级管理者
            if (!permiList.includes(bv) && !isSuperAdmin) {
                // 对没有权限的dom 进行 remove
                return el.parentNode.removeChild(el);
            }
            return;
        },
    })
}