import os, json
from anki.storage import Collection

CARD_MODEL = 'French Basic with Reverse'
DECK_NAME = 'franćais::practice'

##############
# ANKI SETUP #
##############

# get the anki dir and collection database
user_home = os.path.expanduser("~")
anki_home = os.path.join(user_home, '.local/share/Anki2/User 1/collection.anki2')
col = Collection(anki_home, log=True)

# grab the card model and deck
model = col.models.byName(CARD_MODEL)
deck = col.decks.byName(DECK_NAME)

# set the current working deck in the collection
col.decks.select(deck['id'])
col.decks.current()['mid'] = model['id']

#############
# READ DATA #
#############
# with open('imports.json') as f:
#     data = json.load(f)
#     print(data)

# create a new card
note = col.newNote()
# set model deck id to the deck we want to store the note
note.model()['did'] = deck['id']

col.addNote(note)

############################
# SAVE TRANSACTION TO ANKI #
############################
col.save()
