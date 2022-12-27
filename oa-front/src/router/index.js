import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      component: () => import('../views/login/index.vue'),
      meta: { login_require: false, title: "中研环科-管理系统-登录" }
    },
    {
      path: '/',
      component: () => import('../layout/index.vue'),
      redirect: '/login',
      children: [
        {
          path: 'home',
          component: () => import('../views/home/index.vue'),
          meta: { login_require: true, title: "中研环科-管理系统-首页" },
        },
        {
          path: 'expense',
          component: () => import('../views/expense/index.vue'),
          meta: { login_require: true, title: "中研环科-管理系统-我的报销" },
        },
        {
          path: 'contract',
          component: () => import('../views/contract/index.vue'),
          meta: { login_require: true, title: "中研环科-管理系统-合同管理" },
        },
        {
          path: 'pre-research',
          component: () => import('../views/preResearch/index.vue'),
          meta: { login_require: true, title: "中研环科-管理系统-预设计管理" },
        },
        {
          path: 'customer',
          component: () => import('@/views/customer/index.vue'),
          meta: { login_require: true, title: "中研环科-管理系统-客户管理" },
        },
        {
          path: 'product',
          component: () => import('@/views/product/index.vue'),
          meta: { login_require: true, title: "中研环科-管理系统-产品管理" },
        },
        {
          path: 'finance',
          component: () => import('@/views/finance/index.vue'),
          meta: { login_require: true, title: "中研环科-管理系统-财务管理" },
        },
        {
          path: 'enterprise',
          component: () => import('@/views/enterprise/index.vue'),
          meta: { login_require: true, title: "中研环科-管理系统-企业管理" },
        },
        {
          path: 'log',
          component: () => import('@/views/log/index.vue'),
          meta: { login_require: true, title: "中研环科-日志系统" },
        },
      ]
    },
    {
      path: '/:pathMatch(.*)*',
      redirect: '/home',
    }
  ]
})

router.beforeEach((to, from, next) => {
  document.title = to.meta.title
  if (to.meta.login_require) {
    if (localStorage.getItem("token")) {
      next()
    } else {
      next('/')
    }
  } else {
    if (to.path == "/login" && localStorage.getItem("token")) {
      next("/index")
    } else {
      next()
    }
  }
})

export default router
