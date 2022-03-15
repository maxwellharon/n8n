export default [{
        path: '/error-404',
        name: 'error-404',
        component: () =>
            import ('@/views/error/Error404.vue'),
        meta: {
            layout: 'full',
            resource: 'Auth',
            action: 'read',
        },
    },
    {
        path: '/login',
        name: 'auth-login',
        component: () =>
            import ('@/views/pages/authentication/Login-v1.vue'),
        meta: {
            layout: 'full',
            resource: 'Auth',
            redirectIfLoggedIn: true,
        },
    },
    {
        path: '/google-auth',
        name: 'google-auth',
        component: () =>
            import ('@/views/pages/authentication/GoogleAuth.vue'),
        meta: {
            layout: 'full',

        }
    },
    // {
    //   path: '/pages/authentication/login-v1',
    //   name: 'auth-login-v1',
    //   component: () => import('@/views/pages/authentication/Login-v1.vue'),
    //   meta: {
    //     layout: 'full',
    //   },
    // },
    {
        path: '/pages/authentication/login-v2',
        name: 'auth-login-v2',
        component: () =>
            import ('@/views/pages/authentication/Login-v2.vue'),
        meta: {
            layout: 'full',
        },
    },
    {
        path: '/register',
        name: 'auth-register',
        component: () =>
            import ('@/views/pages/authentication/Register.vue'),
        meta: {
            layout: 'full',
            resource: 'Auth',
            redirectIfLoggedIn: true,
        },
    },
    {
        path: '/pages/authentication/register-v1',
        name: 'auth-register-v1',
        component: () =>
            import ('@/views/pages/authentication/Register-v1.vue'),
        meta: {
            layout: 'full',
        },
    },
    {
        path: '/pages/authentication/register-v2',
        name: 'auth-register-v2',
        component: () =>
            import ('@/views/pages/authentication/Register-v2.vue'),
        meta: {
            layout: 'full',
        },
    },
    {
        path: '/forgot-password',
        name: 'auth-forgot-password',
        component: () =>
            import ('@/views/pages/authentication/ForgotPassword.vue'),
        meta: {
            layout: 'full',
            resource: 'Auth',
            redirectIfLoggedIn: true,
        },
    },
    {
        path: '/pages/authentication/forgot-password-v1',
        name: 'auth-forgot-password-v1',
        component: () =>
            import ('@/views/pages/authentication/ForgotPassword-v1.vue'),
        meta: {
            layout: 'full',
        },
    },
    {
        path: '/pages/authentication/forgot-password-v2',
        name: 'auth-forgot-password-v2',
        component: () =>
            import ('@/views/pages/authentication/ForgotPassword-v2.vue'),
        meta: {
            layout: 'full',
            resource: 'Auth',
        },
    },
    {
        path: '/pages/authentication/reset-password-v1',
        name: 'auth-reset-password-v1',
        component: () =>
            import ('@/views/pages/authentication/ResetPassword-v1.vue'),
        meta: {
            layout: 'full',
        },
    },
    {
        path: '/pages/authentication/reset-password-v2',
        name: 'auth-reset-password-v2',
        component: () =>
            import ('@/views/pages/authentication/ResetPassword-v2.vue'),
        meta: {
            layout: 'full',
        },
    },
    {
        path: '/pages/miscellaneous/coming-soon',
        name: 'misc-coming-soon',
        component: () =>
            import ('@/views/pages/miscellaneous/ComingSoon.vue'),
        meta: {
            layout: 'full',
        },
    },
    {
        path: '/pages/miscellaneous/not-authorized',
        name: 'misc-not-authorized',
        component: () =>
            import ('@/views/pages/miscellaneous/NotAuthorized.vue'),
        meta: {
            layout: 'full',
            resource: 'Auth',
        },
    },
    {
        path: '/pages/miscellaneous/under-maintenance',
        name: 'misc-under-maintenance',
        component: () =>
            import ('@/views/pages/miscellaneous/UnderMaintenance.vue'),
        meta: {
            layout: 'full',
        },
    },
    {
        path: '/pages/miscellaneous/error',
        name: 'misc-error',
        component: () =>
            import ('@/views/pages/miscellaneous/Error.vue'),
        meta: {
            layout: 'full',
        },
    },
    {
        path: '/pages/account-setting',
        name: 'pages-account-setting',
        component: () =>
            import ('@/views/pages/account-setting/AccountSetting.vue'),
        meta: {
            pageTitle: 'Account Settings',
            breadcrumb: [{
                    text: 'Pages',
                },
                {
                    text: 'Account Settings',
                    active: true,
                },
            ],
        },
    },
    {
        path: '/pages/settings/workflow',
        name: 'work-flow',
        component: () =>
            import ('@/views/pages/settings/WorkFlow.vue'),
        meta: {
            pageTitle: 'Work-Flow',
            breadcrumb: [{
                    text: 'Work-Flow',
                },
                {
                    text: 'Work-Flow',
                    active: true,
                },
            ],
        },
    },
    {
        path: '/pages/settings/workfloweditor',
        name: 'work-flow-editor',
        component: () =>
            import ('@/views/pages/settings/WorkFlowEditor.vue'),
        meta: {
            pageTitle: 'Work-Flow-Editor',
            breadcrumb: [{
                    text: 'Work-Flow-Editor',
                },
                {
                    text: 'Work-Flow-Editor',
                    active: true,
                },
            ],
        },
    },
    {
        path: '/pages/profile',
        name: 'pages-profile',
        component: () =>
            import ('@/views/pages/profile/Profile.vue'),
        meta: {
            pageTitle: 'Profile',
            breadcrumb: [{
                    text: 'Pages',
                },
                {
                    text: 'Profile',
                    active: true,
                },
            ],
        },
    },
    {
        path: '/pages/profile/profile-signature',
        name: 'pages-signature',
        component: () =>
            import ('@/views/pages/profile/ProfileSignature.vue'),
        meta: {
            pageTitle: 'Profile',
            breadcrumb: [{
                    text: 'Pages',
                },
                {
                    text: 'Profile',
                    active: true,
                },
            ],
        },
    },
    {
        path: '/pages/profile/profile-settings',
        name: 'pages-profile',
        component: () =>
            import ('@/views/pages/profile/ProfileSettings.vue'),
        meta: {
            pageTitle: 'Profile',
            breadcrumb: [{
                    text: 'Pages',
                },
                {
                    text: 'Profile',
                    active: true,
                },
            ],
        },
    },
    {
        path: '/pages/faq',
        name: 'pages-faq',
        component: () =>
            import ('@/views/pages/faq/Faq.vue'),
        meta: {
            pageTitle: 'FAQ',
            breadcrumb: [{
                    text: 'Pages',
                },
                {
                    text: 'FAQ',
                    active: true,
                },
            ],
        },
    },
    {
        path: '/pages/knowledge-base',
        name: 'pages-knowledge-base',
        component: () =>
            import ('@/views/pages/Knowledge-base/KnowledgeBase.vue'),
        meta: {
            pageTitle: 'Knowledge Base',
            breadcrumb: [{
                    text: 'Pages',
                },
                {
                    text: 'Knowledge Base',
                    active: true,
                },
            ],
        },
    },
    {
        path: '/pages/knowledge-base/:category',
        name: 'pages-knowledge-base-category',
        component: () =>
            import ('@/views/pages/Knowledge-base/KnowledgeBaseCategory.vue'),
        meta: {
            pageTitle: 'Category',
            breadcrumb: [{
                    text: 'Pages',
                },
                {
                    text: 'Knowledge Base',
                    to: '/pages/Knowledge-base',
                },
                {
                    text: 'Category',
                    active: true,
                },
            ],
            navActiveLink: 'pages-knowledge-base',
        },
    },
    {
        path: '/pages/knowledge-base/:category/:slug',
        name: 'pages-knowledge-base-question',
        component: () =>
            import ('@/views/pages/Knowledge-base/KnowledgeBaseCategoryQuestion.vue'),
        meta: {
            pageTitle: 'Question',
            breadcrumb: [{
                    text: 'Pages',
                },
                {
                    text: 'Knowledge Base',
                    to: '/pages/Knowledge-base',
                },
                {
                    text: 'Category',
                    to: '/pages/Knowledge-base/category',
                },
                {
                    text: 'Question',
                    active: true,
                },
            ],
            navActiveLink: 'pages-knowledge-base',
        },
    },
    {
        path: '/pages/pricing',
        name: 'pages-pricing',
        component: () =>
            import ('@/views/pages/pricing/Pricing.vue'),
    },
    {
        path: '/pages/blog/list',
        name: 'pages-blog-list',
        component: () =>
            import ('@/views/pages/blog/BlogList.vue'),
        meta: {
            pageTitle: 'Blog List',
            breadcrumb: [{
                    text: 'Pages',
                },
                {
                    text: 'Blog',
                },
                {
                    text: 'List',
                    active: true,
                },
            ],
        },
    },
    {
        path: '/pages/blog/:id',
        name: 'pages-blog-detail',
        component: () =>
            import ('@/views/pages/blog/BlogDetail.vue'),
        meta: {
            pageTitle: 'Blog Detail',
            breadcrumb: [{
                    text: 'Pages',
                },
                {
                    text: 'Blog',
                },
                {
                    text: 'Detail',
                    active: true,
                },
            ],
        },
    },
    {
        path: '/pages/blog/edit/:id',
        name: 'pages-blog-edit',
        component: () =>
            import ('@/views/pages/blog/BlogEdit.vue'),
        meta: {
            pageTitle: 'Blog Edit',
            breadcrumb: [{
                    text: 'Pages',
                },
                {
                    text: 'Blog',
                },
                {
                    text: 'Edit',
                    active: true,
                },
            ],
        },
    },


]