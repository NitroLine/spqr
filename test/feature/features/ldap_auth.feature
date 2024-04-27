Feature: LDAP auth test

  Scenario: Frontend auth works
    Given cluster environment is
    """
    ROUTER_CONFIG=/spqr/test/feature/conf/router_with_ldap_frontend.yaml
    """
    Given cluster is up and running
    When I run command on host "router"
    """
    PGPASSWORD=password psql -c "SELECT 1" -d regress -U regress -p 6432 -h localhost
    """
    Then command return code should be "0"
    And command output should match regexp
    """
    1
    """