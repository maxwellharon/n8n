export default [{
        header: 'Settings',
    },
    {
        title: 'Settings',
        icon: 'SettingsIcon',
        route: null,
    },
    {
        title: 'Workflows',
        icon: 'FastForwardIcon',
        children: [{
                title: 'List',
                route: 'work-flow',
            },
            {
                title: 'New',
                route: null,
            },
        ]

    },
]