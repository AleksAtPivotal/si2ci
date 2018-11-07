BEGIN;
  CREATE OR REPLACE FUNCTION on_pipeline_delete() RETURNS TRIGGER AS $$
  BEGIN
          EXECUTE format('DROP TABLE IF EXISTS pipeline_build_events_%s', OLD.id);
          RETURN NULL;
  END;
  $$ LANGUAGE plpgsql;


  CREATE OR REPLACE FUNCTION on_pipeline_insert() RETURNS TRIGGER AS $$
  BEGIN
          EXECUTE format('CREATE TABLE IF NOT EXISTS pipeline_build_events_%s () INHERITS (build_events)', NEW.id);
          EXECUTE format('CREATE INDEX IF NOT EXISTS pipeline_build_events_%s_build_id ON pipeline_build_events_%s (build_id)', NEW.id, NEW.id);
          EXECUTE format('CREATE UNIQUE INDEX IF NOT EXISTS pipeline_build_events_%s_build_id_event_id ON pipeline_build_events_%s (build_id, event_id)', NEW.id, NEW.id);
          RETURN NULL;
  END;
  $$ LANGUAGE plpgsql;


  DROP TRIGGER IF EXISTS pipeline_build_events_delete_trigger ON pipelines;
  CREATE TRIGGER pipeline_build_events_delete_trigger AFTER DELETE on pipelines FOR EACH ROW EXECUTE PROCEDURE on_pipeline_delete();

  DROP TRIGGER IF EXISTS pipeline_build_events_insert_trigger ON pipelines;
  CREATE TRIGGER pipeline_build_events_insert_trigger AFTER INSERT on pipelines FOR EACH ROW EXECUTE PROCEDURE on_pipeline_insert();
COMMIT;
