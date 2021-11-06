package network

import (
	"l2go-concept/internal/common"
	"l2go-concept/internal/game"
	"l2go-concept/internal/game/model"
)

func UserInfo(client *Client) *common.Buffer {
	player := client.player
	buffer := common.NewBuffer()
	buffer.WriteC(0x04)
	buffer.WriteSD(player.X)
	buffer.WriteSD(player.Y)
	buffer.WriteSD(player.Z)
	buffer.WriteD(0x00) // Boat

	buffer.WriteD(player.EntityId)
	buffer.WriteS(player.Name)
	buffer.WriteD(player.Race)
	buffer.WriteD(player.Sex)

	buffer.WriteD(player.ClassId)

	buffer.WriteD(player.Level)
	buffer.WriteD(player.Experience)

	buffer.WriteD(player.STR)
	buffer.WriteD(player.DEX)
	buffer.WriteD(player.CON)
	buffer.WriteD(player.INT)
	buffer.WriteD(player.WIT)
	buffer.WriteD(player.MEN)

	buffer.WriteD(uint32(player.MaxHp))
	buffer.WriteD(uint32(player.CurrentHp))
	buffer.WriteD(uint32(player.MaxMp))
	buffer.WriteD(uint32(player.CurrentMp))

	buffer.WriteD(player.SkillPoints)
	buffer.WriteD(0x20)  // TODO Current wieght
	buffer.WriteD(0x100) // TODO Max Weight

	buffer.WriteD(0x20) // 20 no weapon, 40 weapon equipped

	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_UNDER))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_REAR))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_LEAR))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_NECK))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_RFINGER))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_LFINGER))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_HEAD))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_RHAND))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_LHAND))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_GLOVES))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_CHEST))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_LEGS))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_FEET))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_BACK))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_LRHAND))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_HAIR))

	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_UNDER))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_REAR))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_LEAR))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_NECK))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_RFINGER))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_LFINGER))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_HEAD))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_RHAND))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_LHAND))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_GLOVES))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_CHEST))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_LEGS))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_FEET))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_BACK))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_LRHAND))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_HAIR))

	buffer.WriteD(1234) // buffer.WriteD(player.PAtk(null))
	buffer.WriteD(600)  // buffer.WriteD(player.PAtkSpd)
	buffer.WriteD(1234) // buffer.WriteD(player.PDef(null))
	buffer.WriteD(1234) // buffer.WriteD(player.EvasionRate(null))
	buffer.WriteD(1234) // buffer.WriteD(player.Accuracy)
	buffer.WriteD(1234) // buffer.WriteD(player.CriticalHit(null, null))
	buffer.WriteD(1234) // buffer.WriteD(player.MAtk(null, null))

	buffer.WriteD(600) // buffer.WriteD(player.MAtkSpd)
	buffer.WriteD(600) // buffer.WriteD(player.PAtkSpd)

	buffer.WriteD(1234) // buffer.WriteD(player.MDef(null, null))

	buffer.WriteD(0x00) // 0-non-pvp 1-pvp = violett name
	buffer.WriteD(player.Karma)

	buffer.WriteD(321) // base run speed
	buffer.WriteD(123) // base walk speed
	buffer.WriteD(100) // swim run speed (calculated by getter)
	buffer.WriteD(80)  // swim walk speed (calculated by getter)

	buffer.WriteD(0x00)
	buffer.WriteD(0x00)

	buffer.WriteD(600) // _flyRunSpd
	buffer.WriteD(400) // _flyWalkSpd

	buffer.WriteF(0.18) // run speed multiplier
	buffer.WriteF(0.18) // attack speed multiplier

	buffer.WriteF(8.0)  // CollisionRadius
	buffer.WriteF(24.0) // CollisionHeight

	buffer.WriteD(player.Hair)
	buffer.WriteD(player.HairColor)
	buffer.WriteD(player.Face)
	buffer.WriteD(0x01) // builder level 1 = GM

	buffer.WriteS(player.Title)

	buffer.WriteD(player.ClanId)
	buffer.WriteD(0x00) //buffer.WriteD(player.ClanCrestId)
	buffer.WriteD(0x00) //buffer.WriteD(player.AllyId)
	buffer.WriteD(0x00) //buffer.WriteD(player.AllyCrestId) // ally crest id

	// 0x40 leader rights
	// siege flags: attacker - 0x180 sword over name, defender - 0x80 shield, 0xC0 crown (|leader), 0x1C0 flag (|leader)
	buffer.WriteD(0x00) // buffer.WriteD(_relation)

	buffer.WriteC(0x00) //buffer.WriteC(player.MountType) // mount type
	buffer.WriteC(0x00) //buffer.WriteC(player.PrivateStoreType)
	buffer.WriteC(0x00) //buffer.WriteC(_player.hasDwarvenCraft() ? 1 : 0)

	buffer.WriteD(player.PkKills)
	buffer.WriteD(player.PvpKills)

	buffer.WriteH(0x00) // buffer.WriteH(player.Cubics().size)
	//for int cubicId : player.Cubics().keySet())
	//{
	//	buffer.WriteH(cubicId);
	//}

	buffer.WriteC(0x00) // buffer.WriteC(_player.isInPartyMatchRoom() ? 1 : 0)

	buffer.WriteD(0x00) // C2 - abnormal effects
	//if player.Appearance().isInvisible()
	//{
	//	buffer.WriteD(player.AbnormalEffect() | Creature.ABNORMAL_EFFECT_STEALTH)
	//}
	//else
	//{
	//	buffer.WriteD(player.AbnormalEffect) // C2
	//}

	buffer.WriteC(0x00) // nothing

	buffer.WriteD(0x00) // buffer.WriteD(player.ClanPrivileges)

	// C4 addition
	buffer.WriteD(0x00) // swim?
	buffer.WriteD(0x00)
	buffer.WriteD(0x00)
	buffer.WriteD(0x00)
	buffer.WriteD(0x00)
	buffer.WriteD(0x00)
	buffer.WriteD(0x00)
	// C4 addition end

	buffer.WriteH(5)    // c2 recommendations remaining
	buffer.WriteH(255)  // c2 recommendations received
	buffer.WriteD(0x00) // player.MountNpcId() > 0 ? player.MountNpcId() + 1000000 : 0
	buffer.WriteH(80)   // buffer.WriteH(player.InventoryLimit) - slots

	buffer.WriteD(player.ClassId)
	buffer.WriteD(0x00) // special effects? circles around player...
	buffer.WriteD(player.MaxCp)
	buffer.WriteD(player.CurrentCp)
	buffer.WriteC(0x00) // buffer.WriteC(_player.isMounted() ? 0 : player.EnchantEffect)

	buffer.WriteC(0x00) // team circle around feet 1= Blue, 2 = red

	buffer.WriteD(0x00) // buffer.WriteD(player.ClanCrestLargeId)
	buffer.WriteC(0x00) // 0x01: Noblesse symbol on char menu ctrl+I
	if player.Hero {
		buffer.WriteC(0x01) // 0x01: Hero Aura
	} else {
		buffer.WriteC(0x00) // 0x01: Hero Aura
	}

	buffer.WriteC(0x00)     // buffer.WriteC(_player.isFishing() ? 1 : 0) // Fishing Mode
	buffer.WriteD(0x00)     // buffer.WriteD(player.FishX)                // fishing x
	buffer.WriteD(0x00)     // buffer.WriteD(player.FishY)                // fishing y
	buffer.WriteD(0x00)     // buffer.WriteD(player.FishZ)                // fishing z
	buffer.WriteD(0xFFFFFF) // buffer.WriteD(player.Appearance().getNameColor)

	return buffer
}

func UserShortcutInit(character *model.Character) *common.Buffer {
	buffer := common.NewBuffer()
	buffer.WriteC(0x45)
	buffer.WriteD(uint32(len(character.Shortcuts)))

	for _, shortcut := range character.Shortcuts {
		buffer.WriteD(shortcut.Type)
		buffer.WriteD(shortcut.Slot + (shortcut.Page * 12))
		buffer.WriteD(shortcut.Id)
		if shortcut.Level > -1 {
			buffer.WriteD(uint32(shortcut.Level)) // Get skill level with shortcut id!
		}
		buffer.WriteD(0x01)
	}
	return buffer
}

func UserItemList(_ *model.Character, showWindow bool) *common.Buffer {
	buffer := common.NewBuffer()
	buffer.WriteC(0x1B)

	if showWindow {
		buffer.WriteH(0x01)
	} else {
		buffer.WriteH(0x00)
	}

	buffer.WriteH(0x00) // Length of the items

	//for (ItemInstance temp : _items)
	//{
	//	buffer.WriteH(temp.getItem().getType1());
	//	buffer.WriteD(temp.getObjectId());
	//	buffer.WriteD(temp.getItemId());
	//	buffer.WriteD(temp.getCount());
	//	buffer.WriteH(temp.getItem().getType2());
	//	buffer.WriteH(temp.getCustomType1());
	//	buffer.WriteH(temp.isEquipped() ? 0x01 : 0x00);
	//	buffer.WriteD(temp.getItem().getBodyPart());
	//	buffer.WriteH(temp.getEnchantLevel());
	//	// race tickets
	//	buffer.WriteH(temp.getCustomType2());
	//}
	return buffer
}

func ClientSetTime(tc game.TimeController) *common.Buffer {
	buffer := common.NewBuffer()

	buffer.WriteC(0xEC)
	buffer.WriteD(uint32(tc.GetGameTime()))
	buffer.WriteD(6) // constant to match the server time( this determines the speed of the client clock)

	return buffer
}

func SkillList(skills []model.Skill) *common.Buffer {
	buffer := common.NewBuffer()
	buffer.WriteC(0x58)
	buffer.WriteD(uint32(len(skills)))

	for _, skill := range skills {
		if skill.Passive {
			buffer.WriteD(0x01)
		} else {
			buffer.WriteD(0x00)
		}

		buffer.WriteD(skill.Level)
		buffer.WriteD(skill.Id)
	}
	return buffer
}

func ValidateLocation(character *model.Character) *common.Buffer {
	buffer := common.NewBuffer()
	buffer.WriteC(0x61)
	buffer.WriteD(character.EntityId)
	buffer.WriteSD(character.X)
	buffer.WriteSD(character.Y)
	buffer.WriteSD(character.Z)
	buffer.WriteSD(character.Heading)
	return buffer
}

func FriendList() *common.Buffer {
	buffer := common.NewBuffer()
	buffer.WriteC(0xFA)

	// TODO actually implement it
	return buffer
}

func QuestList() *common.Buffer {
	buffer := common.NewBuffer()
	buffer.WriteC(0x80)
	buffer.WriteH(0x00)

	// TODO actually implement it
	return buffer
}

func HennaInfo(_ *model.Character) *common.Buffer {
	buffer := common.NewBuffer()
	buffer.WriteC(0xE4)

	buffer.WriteC(0x00) // equip INT
	buffer.WriteC(0x00) // equip STR
	buffer.WriteC(0x00) // equip CON
	buffer.WriteC(0x00) // equip MEM
	buffer.WriteC(0x00) // equip DEX
	buffer.WriteC(0x00) // equip WIT

	// Henna slots
	//int classId = _player.getClassId().level();
	//if (classId == 1)
	//{
	//	buffer.WriteD(2);
	//}
	//else if (classId > 1)
	//{
	//	buffer.WriteD(3);
	//}
	//else
	//{
	buffer.WriteD(0)
	//}

	buffer.WriteD(0x00) // size
	//for (int i = 0; i < _count; i++)
	//{
	//	buffer.WriteD(_hennas[i].getSymbolId());
	//	buffer.WriteD(_hennas[i].canBeUsedBy(_player) ? _hennas[i].getSymbolId() : 0);
	//}

	return buffer
}

func ActionFailed() *common.Buffer {
	buffer := common.NewBuffer()
	buffer.WriteC(0x25)
	return buffer
}
