from peewee import *;

db = SqliteDatabase('sync.db');

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

Block.create(height = 0, time = 'now', txs_n = 0, inner_txs_n = 1, txs = 'david');
print("Insert Succeed!")


# print(Block.select().count())
# for b in Block.select():
#     print(b.time)
