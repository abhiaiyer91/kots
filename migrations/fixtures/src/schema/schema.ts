import * as yaml from "js-yaml";
import * as fs from "fs";
import * as escape from "pg-escape";

export class Schema {
  private parsedDoc: any;

  public parse(filename: string) {
    this.parsedDoc = yaml.safeLoad(fs.readFileSync(filename, "utf-8"));
  }

  public generateFixtures(): string[] {
    let statements: string[] = [];
    console.log("Generating users...");
    for (const user of this.parsedDoc.users) {
      statements = statements.concat(this.generateUserFixture(user));
    }

    console.log("Generating clusters...");
    for (const cluster of this.parsedDoc.clusters) {
      statements = statements.concat(this.generateClusterFixture(cluster));
    }

    console.log("Generating watches...");
    for (const watch of this.parsedDoc.watches) {
      statements = statements.concat(this.generateWatchFixture(watch));
    }

    return statements;
  }

  public generateUserFixture(user: any): string[] {
    const statements: string[] = [
      escape(`insert into github_user (username, github_id, avatar_url, email) values (%L, %L::integer, %L, %L)`,
        user.username, ''+user.github_id, user.avatar_url, user.email),
      escape(`insert into ship_user (id, github_id, created_at) values (%L, %L::integer, %L)`, user.id, ''+user.github_id, user.created_at),
    ];

    const metadata = {};
    metadata[user.username] = user.github_id;

    for (const session of user.sessions) {
      statements.push(
        escape(`insert into session (id, user_id, metadata, expire_at) values (%L, %L, %L, %L)`, session, user.id, JSON.stringify(metadata), new Date().toISOString())
      );
    }

    return statements;
  }

  public generateClusterFixture(cluster: any): string[] {
    const statements: string[] = [
      escape(`insert into cluster (id, title, slug, created_at, updated_at, token, cluster_type) values (%L, %L, %L, %L, %L, %L, %L)`,
        cluster.id, cluster.title, cluster.slug, cluster.created_at, cluster.updated_at, cluster.token, cluster.cluster_type),
    ];

    if (cluster.github) {
      statements.push(
        escape(`insert into cluster_github (cluster_id, owner, repo, branch, installation_id) values (%L, %L, %L, %L, %L::integer)`,
          cluster.id, cluster.github.owner, cluster.github.repo, cluster.github.branch, ''+cluster.github.installation_id),
      );
    }

    for (const user of cluster.users) {
      statements.push(
        escape(`insert into user_cluster (user_id, cluster_id) values (%L, %L)`, user, cluster.id)
      );
    }

    return statements;
  }

  public generateWatchFixture(watch: any): string[] {
    const statements: string[] = [];

    const currentSequenceEscapeSequence = watch.current_sequence === null ? "%L" : "%L::integer";
    const currentSequenceValue = watch.current_sequence === null ? "NULL" : ''+watch.current_sequence;

    statements.push(
      escape(`insert into watch (id, current_state, title, icon_uri, created_at, updated_at, slug, parent_watch_id, current_sequence) values (%L, %L, %L, %L, %L, %L, %L, %L, ${currentSequenceEscapeSequence})`,
        watch.id, watch.current_state, watch.title, watch.icon_uri, watch.created_at, watch.updated_at, watch.slug, watch.parent_watch_id, currentSequenceValue)
    );

    if (watch.cluster) {
      statements.push(
        escape(`insert into watch_cluster (watch_id, cluster_id) values (%L, %L)`, watch.id, watch.cluster),
      );
    }
    for (const user of watch.users) {
      statements.push(
        escape(`insert into user_watch (user_id, watch_id) values (%L, %L)`, user, watch.id)
      );
    }

    if (watch.downstream_tokens) {
      for (const downstreamToken of watch.downstream_tokens) {
        statements.push(
          escape(`insert into watch_downstream_token (watch_id, token) values (%L, %L)`, watch.id, downstreamToken)
        );
      }
    }

    if (watch.versions) {
      for (const version of watch.versions) {
        const pullRequestNumberEscapeSequence = version.pullrequest_number === null ? "%L" : "%L::integer";
        const pullRequestNumberValue = version.pullrequest_number === null ? null : ''+version.pullrequest_number;

        statements.push(
          escape(`insert into watch_version (watch_id, created_at, version_label, status, source_branch, sequence, pullrequest_number) values (%L, %L, %L, %L, %L, %L::integer, ${pullRequestNumberEscapeSequence})`,
            watch.id, version.created_at, version.version_label, version.status, version.source_branch, ''+version.sequence, pullRequestNumberValue),
          escape(`insert into ship_output_files (watch_id, created_at, sequence, filepath) values (%L, %L, %L, %L)`,
            watch.id, version.created_at, ''+version.sequence, `${watch.id}/${version.sequence}.tar.gz`),
          escape(`insert into object_store (filepath, encoded_block) values (%L, %L)`,
            `${watch.id}/${version.sequence}.tar.gz`, version.output),
        );
      }
    }

    return statements;
  }
}
