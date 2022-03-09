export default [{
        header: 'Management',
    },
    {
        title: 'Targets',
        icon: 'TargetIcon',
        // acl: {
        action: 'read',
        resource: 'ACL',
        // },
        children: [{
                title: 'Products',
                route: null,
            },
            {
                title: 'Users',
                route: null,
            },
            {
                title: 'Agencies',
                route: null,
            },

        ],
    },
    {
        title: 'Users',
        icon: 'UsersIcon',
        children: [{
                title: 'List',
                route: 'apps-users-list',
                icon: 'PlusCircleIcon',
            },
            {
                title: 'Teams',
                route: null,
            },
            {
                title: 'Roles',
                route: null,
            },
            {
                title: 'Permissions',
                route: null,
            },

        ],
    },
    {
        title: 'Configuration',
        icon: 'SettingsIcon',
        children: [{
                title: 'Actions Taken',
                route: null,
            },
            {
                title: 'Next Actions',
                route: null,
            },
            {
                title: 'Contact Types',
                route: null,
            },
            {
                title: 'Contact Statuses',
                route: null,
            },
            {
                title: 'Collection Stages',
                route: null,
            },
            {
                title: 'Collection Sub Stages',
                route: null,
            },
            {
                title: 'Debt Types',
                route: null,
            },
            {
                title: 'Buckets',
                route: null,
            },
            {
                title: 'Call Types',
                route: null,
            },
            {
                title: 'Closure Reasons',
                route: null,
            },
            {
                title: 'Dispute Reasons',
                route: null,
            },
            {
                title: 'Delinquency Reasons',
                route: null,
            },

        ],
    },
    {
        title: 'Downloads',
        icon: 'DownloadIcon',
        children: [{
            title: 'Templates',
            route: null,
        }],
    },
    {
        title: 'Company',
        icon: 'BriefcaseIcon',
        children: [{
                title: 'Setup',
                route: null,
            },
            {
                title: 'Branches',
                route: null,
            },
            {
                title: 'Currencies',
                route: null,
            },
            {
                title: 'Countries',
                route: null,
            },

        ],
    },
]