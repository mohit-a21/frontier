import { enforcerContainer } from '../../lib/casbin';

const setupPolicies = async () => {
  await enforcerContainer.enforcer.addJsonPolicy(
    { user: 'alice' },
    {
      entity: 'gojek',
      landscape: 'id',
      environment: 'production',
      team: 'transport'
    },
    { role: 'resource.manager' }
  );

  await enforcerContainer.enforcer.addJsonPolicy(
    { user: 'alice' },
    {
      entity: 'gojek'
    },
    { role: 'dwh.manager' }
  );

  await enforcerContainer.enforcer.addJsonPolicy(
    { user: 'frank' },
    {
      entity: 'gojek',
      landscape: 'id',
      environment: 'production',
      team: 'augur'
    },
    { role: 'resource.manager' }
  );

  await enforcerContainer.enforcer.addJsonPolicy(
    { user: 'bob' },
    {
      team: 'transport'
    },
    { role: 'team.admin' }
  );

  await enforcerContainer.enforcer.addJsonPolicy(
    { user: 'cathy' },
    {
      entity: 'gojek'
    },
    { role: 'entity.admin' }
  );

  await enforcerContainer.enforcer.addJsonPolicy(
    { team: 'transport' },
    {
      team: 'transport'
    },
    { role: 'resource.viewer' }
  );

  await enforcerContainer.enforcer.addJsonPolicy(
    { team: 'augur' },
    {
      team: 'augur'
    },
    { role: 'resource.viewer' }
  );

  await enforcerContainer.enforcer.addJsonPolicy(
    { team: 'gofinance' },
    {
      team: 'gofinance'
    },
    { role: 'resource.viewer' }
  );

  await enforcerContainer.enforcer.addJsonPolicy(
    { team: 'transport' },
    {
      entity: 'gojek',
      privacy: 'public'
    },
    { role: 'resource.viewer' }
  );

  await enforcerContainer.enforcer.addJsonPolicy(
    { team: 'augur' },
    {
      entity: 'gojek',
      privacy: 'public'
    },
    { role: 'resource.viewer' }
  );

  await enforcerContainer.enforcer.addJsonPolicy(
    { team: 'gofinance' },
    {
      entity: 'gofin',
      privacy: 'public'
    },
    { role: 'resource.viewer' }
  );

  await enforcerContainer.enforcer.addPolicy(
    JSON.stringify({ team: 'de' }),
    '*',
    JSON.stringify({ role: 'super.admin' })
  );
};

const setupUserTeamMapping = async () => {
  await enforcerContainer.enforcer.addSubjectGroupingJsonPolicy(
    { user: 'alice' },
    { team: 'transport' }
  );
  await enforcerContainer.enforcer.addSubjectGroupingJsonPolicy(
    { user: 'bob' },
    { team: 'transport' }
  );
  await enforcerContainer.enforcer.addSubjectGroupingJsonPolicy(
    { user: 'dave' },
    { team: 'augur' }
  );
  await enforcerContainer.enforcer.addSubjectGroupingJsonPolicy(
    { user: 'frank' },
    { team: 'augur' }
  );
  await enforcerContainer.enforcer.addSubjectGroupingJsonPolicy(
    { user: 'ele' },
    { team: 'gofinance' }
  );
  await enforcerContainer.enforcer.addSubjectGroupingJsonPolicy(
    { user: 'gary' },
    { team: 'de' }
  );
};

const setupResourceProjectMapping = async () => {
  await enforcerContainer.enforcer.addResourceGroupingJsonPolicy(
    {
      resource: 'p-gojek-id-firehose-transport-123'
    },
    {
      entity: 'gojek',
      environment: 'production',
      landscape: 'id',
      team: 'transport',
      privacy: 'public'
    }
  );

  await enforcerContainer.enforcer.addResourceGroupingJsonPolicy(
    {
      resource: 'p-gojek-id-firehose-augur-345'
    },
    {
      entity: 'gojek',
      environment: 'production',
      landscape: 'id',
      team: 'augur',
      privacy: 'public'
    }
  );

  await enforcerContainer.enforcer.addResourceGroupingJsonPolicy(
    {
      resource: 'p-gojek-id-firehose-augur-private-345'
    },
    {
      entity: 'gojek',
      environment: 'production',
      landscape: 'id',
      team: 'augur',
      privacy: 'private'
    }
  );

  await enforcerContainer.enforcer.addResourceGroupingJsonPolicy(
    {
      resource: 'p-gojek-id-beast-123'
    },
    {
      entity: 'gojek',
      environment: 'production',
      landscape: 'id',
      privacy: 'public'
    }
  );

  await enforcerContainer.enforcer.addResourceGroupingJsonPolicy(
    {
      resource: 'p-gofin-id-firehose-gofinance-789'
    },
    {
      entity: 'gofin',
      environment: 'production',
      landscape: 'id',
      privacy: 'public',
      team: 'gofinance'
    }
  );
};

const setupTeamEntityProjectMapping = async () => {
  await enforcerContainer.enforcer.addResourceGroupingJsonPolicy(
    {
      team: 'augur'
    },
    {
      entity: 'gojek'
    }
  );

  await enforcerContainer.enforcer.addResourceGroupingJsonPolicy(
    {
      team: 'transport'
    },
    {
      entity: 'gojek'
    }
  );

  await enforcerContainer.enforcer.addResourceGroupingJsonPolicy(
    {
      team: 'gofinance'
    },
    {
      entity: 'gofin'
    }
  );
};

const setupPermissionRoleMapping = async () => {
  await enforcerContainer.enforcer.addActionGroupingJsonPolicy(
    { action: '*' },
    { role: 'team.admin' }
  );
  await enforcerContainer.enforcer.addActionGroupingJsonPolicy(
    { action: '*' },
    { role: 'entity.admin' }
  );
  await enforcerContainer.enforcer.addActionGroupingJsonPolicy(
    { action: '*' },
    { role: 'super.admin' }
  );

  await enforcerContainer.enforcer.addActionGroupingJsonPolicy(
    { action: 'firehose.read' },
    { role: 'resource.viewer' }
  );
  await enforcerContainer.enforcer.addActionGroupingJsonPolicy(
    { action: 'dagger.read' },
    { role: 'resource.viewer' }
  );
  await enforcerContainer.enforcer.addActionGroupingJsonPolicy(
    { action: 'beast.read' },
    { role: 'resource.viewer' }
  );

  await enforcerContainer.enforcer.addActionGroupingJsonPolicy(
    { action: 'firehose.write' },
    { role: 'resource.manager' }
  );
  await enforcerContainer.enforcer.addActionGroupingJsonPolicy(
    { action: 'dagger.write' },
    { role: 'resource.manager' }
  );
  await enforcerContainer.enforcer.addActionGroupingJsonPolicy(
    { role: 'resource.viewer' },
    { role: 'resource.manager' }
  );

  await enforcerContainer.enforcer.addActionGroupingJsonPolicy(
    { action: 'beast.*' },
    { role: 'dwh.manager' }
  );
};

export default async () => {
  await setupPolicies();
  await setupUserTeamMapping();
  await setupResourceProjectMapping();
  await setupTeamEntityProjectMapping();
  await setupPermissionRoleMapping();
};