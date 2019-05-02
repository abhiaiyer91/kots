import gql from "graphql-tag";

export const createShipOpsClusterRaw = `
  mutation createShipOpsCluster($title: String!) {
    createShipOpsCluster(title: $title) {
      id
      slug
      shipOpsRef {
        token
      }
    }
  }
`;
export const createShipOpsCluster = gql(createShipOpsClusterRaw);

export const createGitOpsClusterRaw = `
  mutation createGitOpsCluster($title: String!, $installationId: Int, $gitOpsRef: GitOpsRefInput) {
    createGitOpsCluster(title: $title, installationId: $installationId, gitOpsRef: $gitOpsRef) {
      id
      slug
    }
  }
`;
export const createGitOpsCluster = gql(createGitOpsClusterRaw);

export const updateClusterRaw = `
  mutation updateCluster($clusterId: String!, $clusterName: String!, $gitOpsRef: GitOpsRefInput) {
    updateCluster(clusterId: $clusterId, clusterName: $clusterName, gitOpsRef: $gitOpsRef) {
      slug
    }
  }
`;
export const updateCluster = gql(updateClusterRaw);

export const deleteClusterRaw = `
  mutation deleteCluster($clusterId: String!) {
    deleteCluster(clusterId: $clusterId)
  }
`;
export const deleteCluster = gql(deleteClusterRaw);
