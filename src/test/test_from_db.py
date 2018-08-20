from peewee import *;

db = SqliteDatabase('sync.db');
# db = PostgresqlDatabase('xorm_test', user = 'mercury', password='', host='0.0.0.0', port=5432)

class Block(Model):
    height = IntegerField();
    time = CharField();
    txs_n = IntegerField();
    inner_txs_n = IntegerField();
    txs = CharField();
     
    class Meta:
        database = db;

db.connect();
db.create_tables([Block]);

print("total query: [", Block.select().count(), "]")
for b in Block.select():
    print(b.id, "time:", b.time)
